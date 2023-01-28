// Looping Over Arrays and Slices
package main

import "fmt"

func main() {
	names := []string{"Paulo", "Jhon", "Aline", "Maria"}

	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}
}
