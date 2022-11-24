package main

import "fmt"

type newVar int

var x newVar

func main() {
	fmt.Println("x: ", x)
	fmt.Printf("type of x: %T\n", x)
	x = 42
	fmt.Println("x: ", 42)
}
