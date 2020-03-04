package main

import "fmt"

func product(nums ...int) { // HL
	prod := 1
	for _, num := range nums {
		prod *= num
	}
	fmt.Println(nums, prod)
}

// START OMIT
func main() {
	nums := []int{2, 3, 4, -1}
	product(nums...) // HL
}

// END OMIT
