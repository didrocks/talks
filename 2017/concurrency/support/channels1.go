// +build OMIT

package main

import "fmt"

// START OMIT
func main() {

	// Declaring and initializing
	messages := make(chan string) // HL

	// Send a value into a channel (from a goroutine)
	go func() {
		messages <- "ping" // HL
	}()

	// Receive from a channel
	msg := <-messages // HL

	fmt.Println(msg)
}

// STOP OMIT
