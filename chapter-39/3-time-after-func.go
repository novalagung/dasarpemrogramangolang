package main

import "fmt"
import "time"

func main() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")
}
