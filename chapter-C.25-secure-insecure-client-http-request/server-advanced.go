package main

import (
	"crypto/tls"
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

func getTlsConfig() *tls.Config {
	certPair1, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}

	tlsConfig := new(tls.Config)
	tlsConfig.NextProtos = []string{"http/1.1"}
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.PreferServerCipherSuites = true

	tlsConfig.Certificates = []tls.Certificate{
		certPair1, /** add other certificates here **/
	}
	tlsConfig.BuildNameToCertificate()

	tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
	tlsConfig.CurvePreferences = []tls.CurveID{
		tls.CurveP521,
		tls.CurveP384,
		tls.CurveP256,
	}
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	}

	return tlsConfig
}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9000"
	server.TLSConfig = getTlsConfig()

	log.Println("Starting server at", server.Addr)
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}
}
