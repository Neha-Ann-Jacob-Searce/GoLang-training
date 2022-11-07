package main

import "fmt"

func main() {
	nos := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range nos {
		if nos[i]%2 == 0 {
			fmt.Println(nos[i], " is even")
		} else {
			fmt.Println(nos[i], " is odd")
		}
	}
}
