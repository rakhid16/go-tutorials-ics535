package main

import "fmt"

// procedure 1
func plus(a, b int) int {
	return a + b
}

// procedure 2
func mult(a, b int) int {
	return a * b
}

func main() {
	plus_result := plus(1, 2)
	fmt.Println("The result is", plus_result)

	mult_result := mult(3, 4)
	fmt.Println("The result is", mult_result)
}
