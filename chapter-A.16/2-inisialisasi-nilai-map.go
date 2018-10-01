package main

import "fmt"

func main() {
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Println(chicken1)
	fmt.Println(chicken2)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)
}
