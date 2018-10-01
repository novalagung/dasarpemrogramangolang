package main

import "fmt"

func main() {
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	// 104 97 108 111

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)
	// halo
}
