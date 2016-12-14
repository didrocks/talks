// +build OMIT

// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"

// START1 OMIT
func main() {

	// Declaring and initializing
	messages := make(chan string)

	messages <- "ping"
	messages <- "ping" // HL

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// STOP1 OMIT

func send() {
	// Declaring and initializing
	messages := make(chan string)
	// START2 OMIT
	<-messages
	// STOP2 OMIT
}
