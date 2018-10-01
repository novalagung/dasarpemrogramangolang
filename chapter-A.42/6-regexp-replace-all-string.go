package main

import "fmt"
import "regexp"

func main() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
	// "potato potato potato"
}
