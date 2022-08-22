package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// robin := person{"Tim", "Drake"} // problems can arise from this if you modify the struct in any way the values won't be set correctly.
	// fmt.Println(robin)

	nightwing := person{firstName: "Richard", lastName: "Grayson"} // more precise way to use a struct
	fmt.Println(nightwing)

	var spoiler person
	spoiler.firstName = "Stephanie"
	spoiler.lastName = "Brown"
	fmt.Println(spoiler)         // one way to print the values of spoiler
	fmt.Printf("%+v\n", spoiler) // the second way to print the values of structs

	jim := person{
		firstName: "Jim",
		lastName:  "Gordon",
		contactInfo: contactInfo{
			email: "jimbog@gcpd.com",
			zip:   13248,
		},
	}

	jim.updateName("James")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
