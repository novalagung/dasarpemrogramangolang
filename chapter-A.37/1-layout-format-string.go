package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {

	fmt.Printf("%b\n", data.age) // 11010

	fmt.Printf("%c\n", 1400) // ո
	fmt.Printf("%c\n", 1235) // ӓ

	fmt.Printf("%d\n", data.age) // 26

	fmt.Printf("%e\n", data.height) // 1.825000e+02
	fmt.Printf("%E\n", data.height) // 1.825000E+02

	fmt.Printf("%f\n", data.height)   // 182.500000
	fmt.Printf("%.9f\n", data.height) // 182.500000000
	fmt.Printf("%.2f\n", data.height) // 182.50
	fmt.Printf("%.f\n", data.height)  // 182

	fmt.Printf("%e\n", 0.123123123123) // 1.231231e-01
	fmt.Printf("%f\n", 0.123123123123) // 0.123123
	fmt.Printf("%g\n", 0.123123123123) // 0.123123123123

	fmt.Printf("%g\n", 0.12)   // 0.12
	fmt.Printf("%.5g\n", 0.12) // 0.12

	fmt.Printf("%o\n", data.age) // 32

	fmt.Printf("%p\n", &data.name) // 0x2081be0c0

	fmt.Printf("%q\n", `" name \ height "`) // "\" name \\\\ height \""

	fmt.Printf("%s\n", data.name) // wick

	fmt.Printf("%t\n", data.isGraduated) // false

	fmt.Printf("%T\n", data.name)        // string
	fmt.Printf("%T\n", data.height)      // float64
	fmt.Printf("%T\n", data.age)         // int32
	fmt.Printf("%T\n", data.isGraduated) // bool
	fmt.Printf("%T\n", data.hobbies)     // []string

	fmt.Printf("%v\n", data)
	// {wick 182.5 26 false [eating sleeping]}

	fmt.Printf("%+v\n", data)
	// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}

	fmt.Printf("%#v\n", data)
	// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string{"eating", "sleeping"}}

	var people = struct {
		name   string
		height float64
	}{
		name:   "wick",
		height: 182.5,
	}
	var d = data.name

	fmt.Printf("%#v\n", people)
	// struct { name string; height float64 }{name:"wick", height:182.5}

	fmt.Printf("%x\n", data.age)                     // 1a
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3]) // 7769636b
	fmt.Printf("%x\n", d)                            // 7769636b
	fmt.Printf("%%\n")                               // %
}
