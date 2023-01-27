// Using an else if Statement
package main

import "fmt"

func main() {
	input := -3
	if input < 0 {
		fmt.Println("Input can't be a negative number:  ", input)
	} else if input%2 == 0 {
		fmt.Println("Even number: ", input)
	} else {
		fmt.Println("Odd number: ", input)
	}
}
