package main

// START OMIT
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix() // get seconds since 1970
	rand.Seed(int64(sec))    // use seconds to seed random number generator
	r := rand.Int()          // generate a random integer

	switch { // HL
	case r%2 == 0:
		fmt.Println("Random number is even")
	default:
		fmt.Println("Random number is odd")
	} // HL
}

// END OMIT
