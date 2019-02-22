package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func main() {
	bob := person{"Bob", "Newport", contactInfo{email: "steve@mail", zipCode: 12}}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	var travis person
	travis.lastName = "added"
	fmt.Printf("%+v\n%+v\n%+v", alex, bob, travis)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			zipCode: 94000,
			email:   ":JIm@email",
		},
	}
	for i := 0; i < 5; i++ {
		fmt.Println()
	}

	fmt.Printf("%+v", jim)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	jimPointer := &jim
	fmt.Print(jimPointer)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	jimPointer.updateName("tod")
	fmt.Printf("%+v", jim)
	jim.updateName("JIIIIIMMY")
	jim.print()
}
