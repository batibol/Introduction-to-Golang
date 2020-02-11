package main

import (
	"fmt"
	"time"
)

func main() {
	// time in playground is a constant, so this really
	// should be run on your machine...
	switch time.Now().Weekday().String() {
	case "Monday", "Friday", "Sunday":
		fmt.Println("It's a 6-letter day")
	default:
		fmt.Println("It's NOT a 6-letter day")
	}
}
