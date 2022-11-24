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
	//fmt.Printf("%+v", jim)
	jim.print()
	//fmt.Println()
	jim.updateFirstName("Pam")
	jim.print()
	//fmt.Println()

	// jimPointer := &jim
	// jimPointer.updateName("Pam")
	// jim.print()

	jim.updateName("Pam") // pointer shortcut
	jim.print()

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	fmt.Println("Slice updating: doesnt require pointers")
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
	fmt.Println()
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
	fmt.Print("Pass by Value: ")
	p.print()
}

func (pointerToPerson *person) updateName(NewFirstName string) {
	(*pointerToPerson).firstName = NewFirstName
	fmt.Print("Pass by reference: ")
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
