package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact contactInfo
	contactInfo
}

func main() {
	// dunphy := person{firstName: "Dunphy", lastName: "Blah"}
	// dunphy := person {"Dunphy", "Blah"}
	// var dunphy person
	// dunphy.firstName = "Dunphy"
	// dunphy.lastName = "Blah"
	// fmt.Println(dunphy)
	// fmt.Printf("%+v", dunphy)
	jim := person{
		firstName: "Jim",
		lastName:  "Dunder Mifflin",
		contactInfo: contactInfo{
			email:   "jim@dundermifflin.com",
			zipCode: 123456,
		},
	}
	fmt.Printf("%+v", jim)
}
