package main

import "fmt"

// START OMIT
func main() {
	var n []int
	fmt.Println(n, len(n), cap(n))
	if n == nil {
		fmt.Println("nil!")
	}
	z := make([]int, 0)
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

}

// END OMIT
