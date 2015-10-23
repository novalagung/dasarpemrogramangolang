package main

import "fmt"

func main() {
	type student struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}

	var s1 = student{}
	fmt.Println(s1)
}
