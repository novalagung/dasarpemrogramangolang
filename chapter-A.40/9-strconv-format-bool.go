package main

import "fmt"
import "strconv"

func main() {
	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str) // 124
}
