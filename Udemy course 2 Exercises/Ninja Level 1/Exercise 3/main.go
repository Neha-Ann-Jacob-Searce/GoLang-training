package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%s is %d years old? %t maybe", y, x, z)
	fmt.Println(s)
}
