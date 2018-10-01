package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
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

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
