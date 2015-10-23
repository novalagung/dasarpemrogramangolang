package main

import "fmt"
import "os"
import "io"

var path = "/Users/novalagung/Documents/temp/test.txt"

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	fmt.Println(string(text))
	checkError(err)
}

func main() {
	readFile()
}
