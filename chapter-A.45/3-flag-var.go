package main

import "flag"
import "fmt"

func main() {
	// cara ke-1
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)

	// cara ke-2
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Println(data2)
}
