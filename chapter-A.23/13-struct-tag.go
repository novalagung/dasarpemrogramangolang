package main

import "fmt"

type person struct {
	name string `tag1`
	age  int    `tag2`
}

func main() {
	var p1 = person{"wick", 21}
	fmt.Println(p1)
}
