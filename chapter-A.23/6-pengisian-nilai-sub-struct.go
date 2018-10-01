package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func main() {
	var p1 = person{name: "wick", age: 21}
	var s1 = student{person: p1, grade: 2}

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
}
