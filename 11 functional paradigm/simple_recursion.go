package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	// The "factorial" function is calling the "factorial" function
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(4))
}
