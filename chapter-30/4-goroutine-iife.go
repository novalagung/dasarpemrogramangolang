package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	go func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}("wick")

	var message = <-messages
	fmt.Println(message)
}
