package main

import "fmt"
import "strconv"

func main() {
	var str = "true"
	var bul, err = strconv.ParseBool(str)

	if err == nil {
		fmt.Println(bul) // true
	}
}
