package main

import "crypto/sha1"
import "fmt"

func main() {
	var text = "this is secret"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
	// f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8
}
