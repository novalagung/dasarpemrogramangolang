package main

import "fmt"
import "strings"

func main() {
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2)
}
