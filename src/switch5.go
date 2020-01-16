package main

import "fmt"

func main() {
	// START OMIT
	num := 1111

	switch {
	case num > 1000:
		fmt.Println("greater than 1000")
		fallthrough // HL

	case num > 100:
		fmt.Println("greater than 100")
		fallthrough // HL

	case num > 10:
		fmt.Println("greater than 10")

	default:
		fmt.Println("less than 10")
	}
	// END OMIT
}
