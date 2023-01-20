// Implementing Shorthand Operators
package main

import "fmt"

func main() {
	count := 5
	count += 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)
	count -= 5
	fmt.Println(count)

	//also work with strings
	name := "Jhon"
	name += " Textor"
	fmt.Println("Hello,", name)
}
