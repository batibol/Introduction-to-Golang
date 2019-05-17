package main

import "fmt"
// START OMIT
import "time"
import "math/rand"

func main() {
    _, _, sec := time.Now().Clock() 
	rand.Seed(int64(sec))
    i := rand.Int31n(3)

    switch i {
    case 0:
        fmt.Println("zero")
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    }
}
// END OMIT
