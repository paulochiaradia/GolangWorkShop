// Using var to Declare Multiple Variables in One Line
package main

import (
	"fmt"
	"time"
)

func getConfig1() (bool, string, time.Time) {
	return true, "info", time.Now()
}

func main() {
	//type only - if you declare the type all the variables needs to be this type
	var start, middle, end float32
	fmt.Println(start, middle, end)

	//Inicial Value mixed type
	var name, left, right, top, bottom = "one", 1, 1.15, 2, 2.5
	fmt.Println(name, left, right, top, bottom)

	//works with functions also
	var Debug5, LogLevel5, startUpTime = getConfig1()
	fmt.Println(Debug5, LogLevel5, startUpTime)
}
