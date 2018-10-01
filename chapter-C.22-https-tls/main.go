package main

import (
	"log"
	"net/http"
)

func StartNonTLSServer() {
	mux := new(http.ServeMux)
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting to https://localhost/")
		http.Redirect(w, r, "https://localhost/", http.StatusTemporaryRedirect)
	}))

	http.ListenAndServe(":80", mux)
}

func main() {
	go StartNonTLSServer()

	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("Server started at :443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux)
	// err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux)
	if err != nil {
		panic(err)
	}
}
