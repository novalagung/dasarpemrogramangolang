package main

import "fmt"

func main() {
	var c int64 = int64('h')
	fmt.Println(c) // 104

	var d string = string(104)
	fmt.Println(d) // h
}
