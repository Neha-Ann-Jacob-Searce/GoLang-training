package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := person{
		firstName: "pritchett",
		lastName:  "dunphy",
		age:       23,
	}
	fmt.Println(p)
	changeMeVersion1(p)
	fmt.Println(p)
	changeMeVersion2(&p)
	fmt.Println(p)
}

func changeMeVersion1(p person) {
	p.firstName = "Modern"
	p.lastName = "Family"
}

func changeMeVersion2(p *person) {
	(*p).firstName = "Modern"
	(*p).lastName = "Family"
}
