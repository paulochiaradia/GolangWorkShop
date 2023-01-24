// Comparing Values
package main

import "fmt"

func main() {
	visits := 15
	fmt.Println("First visit : ", visits == 1)
	fmt.Println("Return visit : ", visits != 1)
	fmt.Println("Silver Member : ", visits >= 10 && visits <= 20)
	fmt.Println("Gold Member : ", visits >= 21 && visits <= 30)
	fmt.Println("Platinum Member : ", visits > 30)

}
