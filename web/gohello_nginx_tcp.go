package main

import (

	"net"
	"net/http"
	"net/http/fcgi"
	"fmt"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world! on Nginx tcp")
}

func main() {
	l,err := net.Listen("tcp", ":9000")
	if err != nil {
		return
	}
	http.HandleFunc("/", mainHandler)
	fcgi.Serve(l,nil)
}