package main

import "fmt"

func main() {
	print(2022, 1999)
}

func print(current int, birth int) {
	for i := birth; i <= current; i++ {
		fmt.Print(i, " ")
	}
}
