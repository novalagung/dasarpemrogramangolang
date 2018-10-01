package main

import "fmt"
import "strings"

func main() {
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1)

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2)
}
