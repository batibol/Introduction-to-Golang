package main

import "fmt"
// START OMIT
import "time"
import "math/rand"

func somenumber() int32 {
	_, _, sec := time.Now().Clock() // get number of seconds after the hour
	rand.Seed(int64(sec))           // set random seed

	return rand.Int31n(30) - 15
}

func main() {
    if num := somenumber(); num < 0 { // HL
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
//END OMIT
