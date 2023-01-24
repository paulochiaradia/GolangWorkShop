// Iota in golang
package main

import "fmt"

func main() {
	const (
		Sunday  = 0
		Monday  = 1
		Tuesday = 2
	)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	const (
		Sunday1 = iota
		Monday1
		Tuesday1
		Wednesday
	)
	fmt.Println(Sunday1)
	fmt.Println(Monday1)
	fmt.Println(Tuesday1)
	fmt.Println(Wednesday)
}
