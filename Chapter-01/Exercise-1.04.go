// Skipping the Type or Value When Declaring Variables
package main

import (
	"fmt"
	"time"
)

var (
	Debug1       bool
	LogLevel1    = "info"
	startUptime1 = time.Now()
)

func main() {
	fmt.Println(Debug1, LogLevel1, startUptime1)
}
