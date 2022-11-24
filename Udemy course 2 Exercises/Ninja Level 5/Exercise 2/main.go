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
		iceCream:  []string{"Vanilla", "Chocolate"},
	}
	pritchett := person{
		firstName: "Mitchell",
		lastName:  "Gloria",
		iceCream:  []string{"Black currant", "Butterscotch"},
	}

	people := make(map[string]person)
	people["Claire"] = dunphy
	people["Gloria"] = pritchett

	fmt.Println("1st person details:")
	printDetails(dunphy)
	fmt.Println("\n2nd person details:")
	printDetails(pritchett)
	fmt.Println()
	for key, _ := range people {
		printDetails(people[key])
	}
}

func printDetails(p person) {
	fmt.Println("Name: ", p.firstName, " ", p.lastName)
	fmt.Print("Ice cream flavours: ")
	for _, flavour := range p.iceCream {
		fmt.Print(flavour, " ")
	}
	fmt.Println()
}
