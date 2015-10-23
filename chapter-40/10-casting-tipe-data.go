package main

import "fmt"

func main() {
	var a float64 = float64(24)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24
}
