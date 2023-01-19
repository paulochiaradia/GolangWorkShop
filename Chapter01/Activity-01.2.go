// Declaring Multiple Variables with a Short Variable Declaration
package main

import (
	"fmt"
	"time"
)

func main() {
	Debug3, LogLevel3, startUptime3 := false, "info", time.Now()
	fmt.Println(Debug3, LogLevel3, startUptime3)

}
