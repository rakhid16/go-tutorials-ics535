package main

import "fmt"

// Function as a parameter
func apply_function(f func(int) int, x int) int {
	return f(x)
}

// Function as a return value
func multiply_by(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// Function assigned to a variable
	add := func(a int) int {
		return a
	}

	// Pass the "add" function as a parameter to another function
	result := apply_function(add, 3)
	fmt.Println("apply_function result:", result)

	double := multiply_by(2)
	fmt.Println("multiply_by:", double(4))
}
