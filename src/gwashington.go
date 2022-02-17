package main

import (
	"github.com/casimpania/person"
	"fmt"
)

func main() {
	man := person.New(
		"George",
		"Washington",
		290,
	)

	fmt.Println("First Name:", man.GetFirstName())
	fmt.Println("Last Name: ", man.GetLastName())
	fmt.Println("Age: ", man.GetAge())
}
