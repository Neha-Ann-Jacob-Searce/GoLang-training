package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	vehicle1 := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "Red",
		},
		fourWheel: false,
	}
	vehicle2 := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "Black",
		},
		luxury: true,
	}

	fmt.Println("Vehicle 1:", vehicle1)
	fmt.Println("Doors: ", vehicle1.vehicle.doors)
	fmt.Println("\nVehicle 2: ", vehicle2)
	fmt.Println("Doors: ", vehicle2.vehicle.doors)
}
