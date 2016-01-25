package main

import (
	"net/http"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"encoding/json"
	"net/http/fcgi"
	"net"
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

func Route(m *web.Mux) http.Handler {
	m.Get("/user/hello/:name", hello)
	return m
}

func main() {
	l,err := net.Listen("tcp", ":9000")
	if err != nil {
		return
	}
	fcgi.Serve(l, Route(goji.DefaultMux))
}