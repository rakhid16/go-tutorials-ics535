package main

import "fmt"

func hello() func() string {
	// Returning the anonymous function
	return func() string {
		return "Hello world!"
	}
}

func main() {
	hello_world := hello()

	fmt.Println(hello_world())
}
