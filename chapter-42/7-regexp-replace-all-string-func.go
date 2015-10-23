package main

import "fmt"
import "regexp"

func main() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"
}
