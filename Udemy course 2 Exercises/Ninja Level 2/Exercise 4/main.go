package main

import "fmt"

func main() {
	var i int = 1
	fmt.Printf("number: %d, binary form: %b, hexadecimal form: %x\n", i, i, i)
	j := i << 1
	fmt.Printf("number: %d, binary form: %b, hexadecimal form: %x\n", j, j, j)
}
