package main

import "fmt"
import "time"

func main() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = date.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	// 2015-09-02T08:00:00+07:00
}
