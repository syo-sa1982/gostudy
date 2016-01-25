package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world! on Nginx tcp is %s", c.URLParams["name"])
}

func main() {
	l,err := net.Listen("tcp", ":9000")
	if err != nil {
		return
	}
//	http.HandleFunc("/", mainHandler)
//	fcgi.Serve(l,nil)
	goji.Get("/hello/:name", hello)
	fcgi.Serve(l, goji.DefaultMux)
}