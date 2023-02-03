// Using break and continue to Control Loops
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("Skip")
			fmt.Println(r)
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			fmt.Println(r)
			break
		}
	}
}
