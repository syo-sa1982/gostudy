package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"path/filepath"
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}
// HTTPリクエストを処理
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
		log.Println(data)
	}

	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		github.New("e87fd3d4443a41a869eb", "f9efbb46a454cdfbea1e713c9c4bf1df90ceb5f8", "http://localhost:8080/auth/callback/github"),
		facebook.New("866458566797843", "c265d1fe8f5bf6eb28066e20a7f25281", "http://localhost:8080/auth/callback/facebook"),
	)

	r := newRoom()


	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string("/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	http.Handle("/room", r)
	// チャット開始
	go r.run()
	// Webサーバを開始
	log.Println("Webサーバー開始。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
