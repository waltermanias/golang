package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	wally := person{lastname: "Manias", firstname: "Walter", contactInfo: contactInfo{email: "wmanias@gmail.com", zipCode: 5186}}

	wally.print()

	wally.updateName("Cristian")
	wally.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstname = name
}
