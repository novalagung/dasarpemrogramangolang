package main

import "fmt"

func main() {
	var fruits = []string{"apple"}
	var aFruits = []string{"watermelon", "pinnaple"}

	var copiedElemen = copy(fruits, aFruits)

	fmt.Println(fruits)       // ["watermelon"]
	fmt.Println(aFruits)      // ["watermelon", "pinnaple"]
	fmt.Println(copiedElemen) // 1
}
