package main

import "fmt"
import "os"

var path = "/Users/novalagung/Documents/temp/test.txt"

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func deleteFile() {
	var err = os.Remove(path)
	checkError(err)
}

func main() {
	deleteFile()
}
