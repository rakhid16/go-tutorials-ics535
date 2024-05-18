// In another package
package main

import (
	"fmt"
	"somepackage/somesubpackage"
)

func main() {
	// Using the exported function
	sum := somesubpackage.Add(2, 3)
	fmt.Println("Sum:", sum)

	// Using the exported struct and method
	calc := somesubpackage.Calculator{}
	calc.Add(1, 2, 3)
	fmt.Println("Result:", calc.Result)
}
