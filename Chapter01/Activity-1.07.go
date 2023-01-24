// Pointer Value Swap
package main

import "fmt"

func main() {
	a, b := 10, 5
	fmt.Printf("A antes123: %#v\n", a)

	swap(&a, &b)
	fmt.Printf("A depois123: %#v\n", a)

	fmt.Println(a == 5, b == 10)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
