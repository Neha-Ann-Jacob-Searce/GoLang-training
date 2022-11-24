package main

import "fmt"

func main() {
	if 10 == 5*2 {
		fmt.Print("5*2=10 ", true)
	} else {
		fmt.Print("5*2=10 ", false)
	}
}
