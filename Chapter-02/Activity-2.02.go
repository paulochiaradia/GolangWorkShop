// Looping Over Map Data Using range
package main

import "fmt"

func main() {
	table := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	topWord := ""
	topCount := 0

	for key, value := range table {
		if value > topCount {
			topWord = key
			topCount = value
		}
	}
	fmt.Println("Most popular word: ", topWord)
	fmt.Println("With a count of: ", topCount)

}
