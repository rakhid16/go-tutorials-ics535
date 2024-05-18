package main

import "fmt"

func main() {
	var test int = 32
	var test1 float64 = float64(test)

	fmt.Println("Original value =", test)
	fmt.Println("Converted value =", test1)
}
