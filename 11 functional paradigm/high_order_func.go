package main

import "fmt"

// apply_function applies the given function f to the
// integer x and returns the result.
func apply_function(f func(int) int, x int) int {
	return f(x)
}

func main() {
	// Define a function that doubles the given integer.
	double := func(x int) int {
		return x * 2
	}

	// Apply the double function to the integer 3.
	result := apply_function(double, 4)
	fmt.Println(result)
}
