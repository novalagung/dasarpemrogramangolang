package main

import "fmt"
import "strconv"

func main() {
	var str = "1010"
	var num, err = strconv.ParseInt(str, 2, 8)

	if err == nil {
		fmt.Println(num) // 10
	}
}
