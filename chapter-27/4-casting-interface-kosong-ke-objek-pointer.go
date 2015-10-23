package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	var secret interface{} = &person{name: "wick", age: 27}
	var name = secret.(*person).name
	fmt.Println(name)
}
