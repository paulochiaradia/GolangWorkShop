// Zero Values
package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		count     int
		discount  float64
		debug     bool
		message   string
		emails    []string
		startTime time.Time
	)
	fmt.Printf("Count : %#v \n", count)
	fmt.Printf("Discount : %#v \n", discount)
	fmt.Printf("debug : %#v \n", debug)
	fmt.Printf("Message : %#v \n", message)
	fmt.Printf("emails : %#v \n", emails)
	fmt.Printf("startTime : %#v \n", startTime)
}
