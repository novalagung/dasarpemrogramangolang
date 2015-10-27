package main

import "fmt"
import "net/http"
import "encoding/json"
import "bytes"
import "net/url"

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUser(ID string) student {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	return data
}

func main() {
	var user1 = fetchUser("E001")
	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}
