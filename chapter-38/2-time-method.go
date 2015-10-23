package main

import "fmt"
import "time"

func main() {
	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())
	// year: 2015 month: 8
}
