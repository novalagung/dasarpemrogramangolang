package main

import "fmt"

func main() {
	var c []byte = []byte("halo")
	fmt.Println(c) // [104 97 108 111]

	var d string = string([]byte{104, 97, 108, 111})
	fmt.Println(d) // halo
}
