package main

import (
	"fmt"
	"time"
)

// Chef function
func chef(name string, taskDone chan<- string) {
	fmt.Println(name, "is cooking...")
	time.Sleep(time.Second * 2) // Simulating cooking time
	taskDone <- name + " has finished cooking!"
}

func main() {
	taskDone := make(chan string, 2) // Buffered channel

	// Start three chef goroutines
	go chef("Chef Anne", taskDone)
	go chef("Chef John", taskDone)
	go chef("Chef Juna", taskDone)

	// Await results from our chefs
	for i := 0; i < 3; i++ {
		fmt.Println(<-taskDone)
	}
}
