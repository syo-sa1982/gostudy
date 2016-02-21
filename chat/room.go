package main
import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"github.com/syo-sa1982/gostudy/trace"
)

type room struct {
	// 転送メッセージ
	forward chan *message
	// 入室チャネル
	join chan *client
	// 退室チャネル
	leave chan *client
	// 在室中
	clients map[*client]bool

	tracer trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer: trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
			r.tracer.Trace("新しいクライアント追加")
		case client := <-r.leave:

			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("クライアント退室")
		case msg := <-r.forward:
			r.tracer.Trace("メッセージ受信: ", msg.Message)
			for client := range r.clients {
				select {
				case client.send <- msg:
					r.tracer.Trace(" -- クライアントに送信されました")
				default:
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- 送信失敗しました。クライアントをクリーンアップ")
				}
			}
		}
	}
}

const (
	socketBufferSize = 1024
	messeageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan *message, messeageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}