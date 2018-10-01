package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type M map[string]interface{}

func ActionData(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request with method", r.Method)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	payload := make(M)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := payload["Name"]; !ok {
		http.Error(w, "Payload `Name` is required", http.StatusBadRequest)
		return
	}

	data := M{
		"Message": fmt.Sprintf("Hello %s", payload["Name"]),
		"Status":  true,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9000"

	log.Println("Starting server at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}
}
