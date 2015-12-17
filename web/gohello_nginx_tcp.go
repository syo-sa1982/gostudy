package main

import (
//	"net"
	"net/http"
//	"net/http/fcgi"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"encoding/json"
)

type User struct {
	Id   int64
	Name string
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	u := User{
		Id:   1,
		Name: c.URLParams["name"],
	}

	j, _ := json.Marshal(u)
	fmt.Fprintf(w, string(j))
}

func Route(m *web.Mux) {
	m.Get("/hello/:name", hello)
}

func main() {
//	l,err := net.Listen("tcp", ":9000")
//	if err != nil {
//		return
//	}
	Route(goji.DefaultMux)
	goji.Serve()
//	fcgi.Serve(l, Route(goji.DefaultMux))
}