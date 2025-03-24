package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// main function
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "kVb7r@example.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// ______________________________________________________________________
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// ______________________________________________________________________
func (p person) print() {
	fmt.Printf("%+v", p)
}
