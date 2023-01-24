// Function Design with Pointers
package main

import "fmt"

func main() {
	count := 5
	addValue(count)
	fmt.Println("addValue post: ", count)
	addPointer(&count)
	fmt.Println("addPointer post: ", count)

}

func addValue(count int) {
	count += 5
	fmt.Println("addValue: ", count)
}

func addPointer(count *int) {
	*count += 5
	fmt.Println("addValue: ", *count)
}
