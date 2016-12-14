// +build OMIT

package main

import "fmt"
import "time"

// START OMIT
func main() {

	// Declaring and initializing
	messages := make(chan string)

	// Send a value into a channel (from a goroutine)
	go func() {
		messages <- "ping"
		fmt.Println(<-messages) // HL
	}()

	// Receive from a channel
	msg := <-messages
	fmt.Println(msg)

	messages <- "pong" // HL
	time.Sleep(time.Second)
}

// STOP OMIT
