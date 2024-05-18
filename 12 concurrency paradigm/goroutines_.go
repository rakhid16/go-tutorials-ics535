package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("From goroutine:", i)
	}
}

func main() {
	start_time := time.Now()

	// Launching a goroutine
	go printNumbers()

	// Main function continues to execute concurrently with the goroutine
	for i := 1; i <= 5; i++ {
		fmt.Println("From main:", i)
		time.Sleep(1 * time.Second)
	}

	end_time := time.Now()

	fmt.Println("Total execution time:", end_time.Sub(start_time))
}
