package main

import "fmt"

// START OMIT
import "regexp"

func main() {
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var shortPhone = regexp.MustCompile(`^[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)
	var longPhone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case shortPhone.MatchString(contact):
		fmt.Println(contact, "is a short phone number")
	case longPhone.MatchString(contact):
		fmt.Println(contact, "is a long phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}
}

//END OMIT
