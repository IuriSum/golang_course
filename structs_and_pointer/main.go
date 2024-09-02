package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{"Nome1", "Nome2", contactInfo{"asdsad", 123}}
	fmt.Println(alex)

	var joaoooo person
	joaoooo.firstName = "Joao"
	joaoooo.lastName = "Pedro"
	fmt.Printf("%v", joaoooo)

	joaoooo.contact = contactInfo{"Endereco", 80215250}
	pointer := &joaoooo //& gives the memory address to variable
	pointer.updateName("TOME PONTEIRO")
	joaoooo.print()

}

func (personPointer *person) updateName(newFirstName string) { // *pointer notation gives the value associated to the memory address
	(*personPointer).firstName = newFirstName // (*personPointer) notation gives the
}

func (p person) print() { //associates function direct to struct
	fmt.Printf("%v", p)
}
