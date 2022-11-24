package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	dunphy := person{
		firstName: "Phil",
		lastName:  "Claire",
		iceCream:  []string{"vanilla", "chocolate"},
	}
	pritchett := person{
		firstName: "Mitchell",
		lastName:  "Gloria",
		iceCream:  []string{"Black Currant", "Butterscotch"},
	}
	fmt.Println("1st person details:")
	printDetails(dunphy)
	fmt.Println("\n2nd person details:")
	printDetails(pritchett)
}

func printDetails(p person) {
	fmt.Println("Name: ", p.firstName, " ", p.lastName)
	fmt.Print("Ice cream flavours: ")
	for _, flavour := range p.iceCream {
		fmt.Print(flavour, " ")
	}
	fmt.Println()
}
