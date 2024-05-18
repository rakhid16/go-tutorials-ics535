package main

import "fmt"

func main() {
	const acc_gravity = 9.8 // Immmutable variable
	var mass = 5            // Mutable variable (not following functional programming)

	gravity_force := acc_gravity * float64(mass)

	fmt.Println(gravity_force)
}
