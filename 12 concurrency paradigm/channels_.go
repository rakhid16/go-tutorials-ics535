package main

import "fmt"

func producer(ch chan<- int) {
	defer close(ch) // Close the channel when producer is done

	for i := 1; i <= 5; i++ {
		fmt.Println("Sending data", i, "to channel")
		ch <- i // Send data to the channel
	}
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num) // Receive data from the channel
	}
}

func main() {
	// Create a channel to communicate between producer and consumer
	ch := make(chan int)

	// Start the producer goroutine
	go producer(ch)

	// Start the consumer goroutine
	go consumer(ch)

	// Wait for both goroutines to finish
	var input string

	fmt.Println("Press any key to exit...\n")
	fmt.Scanln(&input)
}
