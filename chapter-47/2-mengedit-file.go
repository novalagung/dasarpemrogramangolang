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

func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	checkError(err)
	_, err = file.WriteString("mari belajar golang\n")
	checkError(err)

	// simpan perubahan
	err = file.Sync()
	checkError(err)
}

func main() {
	writeFile()
}
