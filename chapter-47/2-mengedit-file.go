package main

import "fmt"
import "os"

var path = "/Users/novalagung/Documents/temp/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func writeFile() {
    // buka file dengan level akses READ & WRITE
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

    // tulis data ke file
    _, err = file.WriteString("halo\n")
	if isError(err) { return }
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) { return }

    // simpan perubahan
    err = file.Sync()
	if isError(err) { return }

	fmt.Println("==> file berhasil di isi")
}

func main() {
	writeFile()
}
