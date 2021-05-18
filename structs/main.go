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
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 94000,
		},
	}
	// not gonna work updating this way, need to use pointers
	// works like "Pass by reference"
	//jim.updateNameOld("Jimmy")

	// use of pointers
	// pass by value
	// give me the memory address the variable (jim) pointing at
	// jim is a refrence to the struct in memory, the actual value of the struct
	// &jim as a whole is memory address or a pointer which we are assigning to jimPointer
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

}

// works as "pass by reference" means
// person struct which come here never gonna change in this function
// func (p person) updateNameOld(newFirstName string) {
// 	p.firstName = newFirstName
// }

// Turn address into value with *address
// Turn value into address with &value

// *person means a pointer that points out at person value
// so it's type Description means we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {

	// give me value this memory address is pointing at
	// *pointerToPerson is an operator, means we want to manipulate the value which
	// the pointer is Referencing
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
