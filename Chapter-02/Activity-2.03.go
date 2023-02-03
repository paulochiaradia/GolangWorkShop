// Bubble Sort
package main

import "fmt"

func main() {
	slice := []int{1, 9, 2, 8, 3, 6, 5, 4, 7}
	fmt.Println("Unsorted:", slice)
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				swapped = true
			}
		}
	}
	fmt.Println("Sorted:", slice)
}
