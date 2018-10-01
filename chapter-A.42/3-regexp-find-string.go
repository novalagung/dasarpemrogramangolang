package main

import "fmt"
import "regexp"

func main() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.FindString(text)
	fmt.Println(str)
	// "banana"
}
