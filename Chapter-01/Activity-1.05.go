package main

import "fmt"

var defaultOffSet = 10

func main() {
	offset := defaultOffSet
	fmt.Println(offset)
	offset = offset + defaultOffSet
	fmt.Println(offset)
}
