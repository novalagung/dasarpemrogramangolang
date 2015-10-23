package main

import "fmt"

func main() {
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}
}
