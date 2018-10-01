package main

import "fmt"
import "strings"

func main() {
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1)

	var string2 = strings.Split("batman", "")
	fmt.Println(string2)
}
