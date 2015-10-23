package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}
