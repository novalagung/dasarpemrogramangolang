package main

import "fmt"
import "strings"

func main() {
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)
}
