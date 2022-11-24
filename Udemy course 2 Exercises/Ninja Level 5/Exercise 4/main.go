package main

import "fmt"

func main() {
	p := struct {
		firstName string
		lastName  string
		iceCream  []string
	}{
		firstName: "Phil",
		lastName:  "Claire",
		iceCream:  []string{"Vanilla", "Chocolate"},
	}

	fmt.Println("Name: ", p.firstName, " ", p.lastName)
	fmt.Print("Ice cream flavours: ")
	for _, flavour := range p.iceCream {
		fmt.Print(flavour, " ")
	}
	fmt.Println()
}
