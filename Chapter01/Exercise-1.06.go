// Declaring Multiple Variables from a Function
package main

import (
	"fmt"
	"time"
)

// func name(parameter types and names) (return types)
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug4, LogLevel4, startUptime4 := getConfig()
	fmt.Println(Debug4, LogLevel4, startUptime4)
}
