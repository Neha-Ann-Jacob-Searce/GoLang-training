package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("Decimal version of %v: %d\n", x, x)
	fmt.Printf("Binary version of %v: %b\n", x, x)
	fmt.Printf("Octal version of %v: %o\n", x, x)
	fmt.Printf("Hexidecimal version of %v: %x\n", x, x)
}
