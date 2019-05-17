package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
    _, _, sec := time.Now().Clock()
	rand.Seed(int64(sec))

    for { // HL
        fmt.Println("Infinite loop!")
		if rand.Int31n(10) > 8 {
        	break
		}
    } // HL
}
