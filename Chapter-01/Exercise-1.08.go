// Changing Multiple Values at Once
package main

import "fmt"

func main() {
	/*var first, second, third float32
	first, second, third = 0, 1, 2
	fmt.Println(first, second, third)
	first, second, third = 3, 4, 5
	fmt.Println(first, second, third)*/
	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)
	query, limit, offset = "men", 20, offset+1
	fmt.Println(query, limit, offset)
}
