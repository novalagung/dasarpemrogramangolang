package main

import "fmt"
import "time"

func main() {
	<-time.After(4 * time.Second)
	fmt.Println("expired")
}
