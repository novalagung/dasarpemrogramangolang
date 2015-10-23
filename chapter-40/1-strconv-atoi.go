package main

import "fmt"
import "strconv"

func main() {
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}
}
