package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix() // seconds since 1970...
	rand.Seed(sec)           // seed the random number generator

	for { // HL
		fmt.Println("Infinite loop!")
		if rand.Int31n(10) > 8 {
			fmt.Println("break!")
			break
		}
	} // HL
}
