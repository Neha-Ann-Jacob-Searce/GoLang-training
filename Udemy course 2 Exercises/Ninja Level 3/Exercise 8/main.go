package main

import (
	"fmt"
	"time"
)

func main() {
	i := time.Now()
	switch {
	case i.Hour() < 12:
		fmt.Println("It's forenoon")
	case i.Hour() > 12 || (i.Hour() == 12 && i.Minute() > 0):
		fmt.Println("It's afternoon")
	case i.Hour() == 12 && i.Minute() == 0:
		fmt.Println("It's noon")
	}
}
