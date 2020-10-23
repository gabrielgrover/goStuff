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

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (pp *person) updateName(newFirstName string) {
	(*pp).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
