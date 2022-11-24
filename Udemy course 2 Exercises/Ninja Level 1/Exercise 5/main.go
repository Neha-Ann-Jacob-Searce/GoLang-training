package main

import "fmt"

type newVar int

var x newVar
var y int

func main() {
	fmt.Println("x: ", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 42
	fmt.Println("x: ", x)
	y = int(x)
	fmt.Println("y: ", y)
	fmt.Printf("Type of y: %T\n", y)
}
