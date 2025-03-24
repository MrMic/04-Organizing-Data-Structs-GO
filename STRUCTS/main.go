package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// main function
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "kVb7r@example.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim) // %+v is a verb to print a ...any)
}
