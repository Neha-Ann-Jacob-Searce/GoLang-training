package main

import "fmt"

func main() {
	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	slice := [][]string{slice1, slice2}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Println(slice[i][j])
		}
		fmt.Println()
	}
}
