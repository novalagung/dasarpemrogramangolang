package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}
