package main

import "fmt"

type Mountain struct {
	Name      string
	Elevation int
}

//START OMIT
type Climb struct {
	Mountain
	Climber string
}

func (m Mountain) HowBig() string {
	if m.Elevation > 8800 {
		return "HUGE"
	}
	return "NORMAL"
}

// This method overrides the one above
func (c Climb) HowBig() string {
	return "OVERRIDE!"
}

func main() {
	c2 := Climb{
		Mountain{"K2", 8111}, "Arjun Climber",
	}
	fmt.Println(c2.HowBig())          // will call overriding method
	fmt.Println(c2.Mountain.HowBig()) // will call overridden method
}

//END OMIT
