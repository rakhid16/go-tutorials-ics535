package main

import "fmt"

type Vehicle struct {
	Seats int
	Color string
}

type Car struct {
	Vehicle
	year int
}

type MotorCycle struct {
	Base Vehicle
}

func main() {
	car := &Car{
		year: 1999,
		Vehicle: Vehicle{
			Seats: 4,
			Color: "blue",
		},
	}

	motorCycle := &MotorCycle{
		Vehicle{
			Seats: 2,
			Color: "red",
		},
	}

	fmt.Println(car.Seats)
	fmt.Println(car.Color)
	fmt.Println(car.year)

	fmt.Println(motorCycle.Base.Seats)
	fmt.Println(motorCycle.Base.Color)
}
