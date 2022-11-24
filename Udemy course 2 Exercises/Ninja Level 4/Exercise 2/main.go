package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, val := range slice {
		fmt.Println(val)
	}
	fmt.Printf("%T\n", slice)
}
