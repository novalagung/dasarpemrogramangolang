package main

import (
	// "golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9000"

	// http2.ConfigureServer(server, nil)

	log.Println("Server started at :9000")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
}
