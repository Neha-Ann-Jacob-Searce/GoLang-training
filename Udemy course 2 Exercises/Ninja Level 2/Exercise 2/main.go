package main

import "fmt"

func main() {
	g := 12 == 13
	h := 14 <= 20
	i := 12 >= 12
	j := 12 != 12
	k := 12 < 12
	l := 13 > 9
	fmt.Println("12 == 13 ", g)
	fmt.Println("14 <= 20 ", h)
	fmt.Println("12 >= 12 ", i)
	fmt.Println("12 != 12 ", j)
	fmt.Println("12 < 12 ", k)
	fmt.Println("13 > 9 ", l)
}
