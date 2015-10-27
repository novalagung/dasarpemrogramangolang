package main

import "fmt"
import "net/http"
import "encoding/json"

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() []*student {
	var err error
	var client = &http.Client{}
	var data []*student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

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
	var users = fetchUsers()

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}
