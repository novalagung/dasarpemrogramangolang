package main

import "fmt"
import "time"

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time1 2015-09-01 17:59:31.73600891 +0700 WIB

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
	// time2 2011-12-24 10:20:00 +0000 UTC
}
