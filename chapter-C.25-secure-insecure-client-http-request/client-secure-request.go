package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type M map[string]interface{}

func doRequest(url, method string, data interface{}) (interface{}, error) {
	var payload *bytes.Buffer = nil

	if data != nil {
		payload = new(bytes.Buffer)
		err := json.NewEncoder(payload).Encode(data)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	certFile, err := ioutil.ReadFile("server.crt")
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certFile)

	tlsConfig := &tls.Config{RootCAs: caCertPool}
	tlsConfig.BuildNameToCertificate()

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	responseBody := make(M)
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func main() {
	baseURL := "https://localhost:9000"
	method := "POST"
	data := M{"Name": "Noval Agung"}

	responseBody, err := doRequest(baseURL+"/data", method, data)
	if err != nil {
		log.Println("ERROR", err.Error())
		return
	}

	log.Printf("%#v \n", responseBody)
}

// ======= you can also use both crt and key using code below

// cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
// if err != nil {
//  log.Fatalln("Unable to load cert", err)
// }

// clientCACert, err := ioutil.ReadFile("server.crt")
// if err != nil {
//  log.Fatal("Unable to open cert", err)
// }

// clientCertPool := x509.NewCertPool()
// clientCertPool.AppendCertsFromPEM(clientCACert)

// tlsConfig := &tls.Config{
//  Certificates: []tls.Certificate{cert},
//  RootCAs:      clientCertPool,
// }
