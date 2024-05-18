package main

import "fmt"

// Define an interface
type Animal interface {
	Sound() string
}

// Define a struct for Dog
type Dog struct {
	Name string
}

// Implement the Sound method for Dog
func (d Dog) Sound() string {
	return "Woof! My name is " + d.Name
}

// Define a struct for Cat
type Cat struct {
	Name string
}

// Implement the Sound method for Cat
func (c Cat) Sound() string {
	return "Meow! My name is " + c.Name
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{"Buddy"}
	cat := Cat{"Whiskers"}

	// Call the Sound method for Dog and Cat
	fmt.Println(dog.Sound())
	fmt.Println(cat.Sound())
}
