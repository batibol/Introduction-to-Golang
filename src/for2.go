package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
    _, _, sec := time.Now().Clock()
	rand.Seed(int64(sec))
    num := rand.Int31n(10) + 5

    for ; num > 0; num-- { // HL
        fmt.Println(num)
    } // HL
    fmt.Println("Blast off!")
}
