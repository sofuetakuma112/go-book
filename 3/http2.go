package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	// GoでHTTPSサーバーを起動する際は、デフォルトでHTTP/2を使う
	server.ListenAndServeTLS("cert.pem", "key.pem") // HTTPSサーバーを起動する
}
