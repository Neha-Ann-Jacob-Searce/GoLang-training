package main

import "fmt"

func main() {
	if 10 == 5*3 {
		fmt.Print("10 = 5*3 ", true)
	} else if 10 == 5*4 {
		fmt.Print("10 = 5*4 ", true)
	} else {
		fmt.Print("10 = 5*4 ", false)
	}
}
