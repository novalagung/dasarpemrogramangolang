package main

import "fmt"
import "strings"

func main() {
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)
}
