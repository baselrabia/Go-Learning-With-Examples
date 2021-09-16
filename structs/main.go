package main

import "fmt"



type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}


func main() {


	alex := person{
		firstName: "alex",
		lastName: "anderson",
		contact: contactInfo{
			email: "alex@gmail.com",
			zip: 94000,
		},
	}


	alex.print()
 

	alex.updateName("jim")
	alex.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}