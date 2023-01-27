// A Simple if Statement
package main

import "fmt"

func main() {
	input := 6
	//odd number

	if input%2 == 0 {
		fmt.Println(input, "is an even number ")
	} else {
		fmt.Println(input, "is an odd number: ")
	}
}
