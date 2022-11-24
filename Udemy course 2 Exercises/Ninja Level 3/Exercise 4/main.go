package main

import "fmt"

func main() {
	print(2022, 1999)
}

func print(current int, birth int) {
	i := birth
	for {
		if i > current {
			break
		} else {
			fmt.Print(i, " ")
		}
		i++
	}
}
