package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//seed random number generator using the current time
	rand.Seed(time.Now().UnixNano())
	//Generate the random number between 0-5 than add 1 to get 1-5
	r := rand.Intn(5) + 1
	//Using the string repeater to creat the string with the correct number of stars
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
