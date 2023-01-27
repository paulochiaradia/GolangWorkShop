// Implementing FizzBuzz
package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*	My code
		for i := 1; i <= 100; i++ {
			if (i%5 == 0) && (i%3 == 0) {
				fmt.Println("FizzBuzz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else if i%3 == 0 {
				fmt.Println("Fizz")
			} else {
				fmt.Println(i)
			}

		}*/
	//Book code - more optimized
	for i := 1; i <= 100; i++ {
		out := ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out == "" {
			out = strconv.Itoa(i)
		}
		fmt.Println(out)
	}

}
