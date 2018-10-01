package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var student struct {
		person
		grade int
	}

	student.person = person{"wick", 21}
	student.grade = 2

	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println(student.grade)
}
