package main

import "fmt"

func main() {
	var favSport string
	fmt.Println("Sport: ")
	fmt.Scan(&favSport)
	switch favSport {
	case "Cricket":
		fmt.Println("Cricket")
	case "Football":
		fmt.Println("Football")
	default:
		fmt.Println(favSport)
	}
}
