package main

import "fmt"

func main() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}
}
