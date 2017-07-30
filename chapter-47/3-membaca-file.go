package main

import "fmt"
import "os"
import "io"

var path = "/Users/novalagung/Documents/temp/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func readFile() {
    // buka file
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

    // baca file
    var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) { break }
		}
		if n == 0 {
			break
		}
	}
	if isError(err) { return }

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func main() {
	readFile()
}
