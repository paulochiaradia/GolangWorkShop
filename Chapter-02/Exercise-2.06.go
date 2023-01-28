// switch Statements and Multiple case Values
package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Tuesday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Sunday, time.Saturday:
		fmt.Println("Born on a Weekend")
	default:
		fmt.Println("Error, day born not valid")
	}

}
