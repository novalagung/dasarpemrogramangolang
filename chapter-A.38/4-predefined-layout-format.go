package main

import "fmt"
import "time"

func main() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date.String())
	// 2015-09-02 08:00:00 +0700 WIB
}
