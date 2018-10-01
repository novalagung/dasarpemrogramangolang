package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s *student) sayHello() {
	fmt.Println("halo", s.name)
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var s2 = &student{"ethan hunt", 22}
	s2.sayHello()
}
