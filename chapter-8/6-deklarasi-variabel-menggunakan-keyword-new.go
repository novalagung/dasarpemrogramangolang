package main

import "fmt"

func main() {
	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""
}
