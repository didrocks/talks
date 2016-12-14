// +build OMIT

package main

import "fmt"
import "time"

func main() {
	// START1 OMIT
	a, b := make(chan bool), make(chan bool)
	go func() {
		close(a)
		time.Sleep(time.Millisecond) // HL
		close(b)
	}()
	WaitMany(a, b) // HL
	// STOP1 OMIT
}

// START2 OMIT
// WaitMany waits for a and b to close.
func WaitMany(a, b chan bool) {
	var aclosed, bclosed bool
	for !aclosed || !bclosed { // HL
		fmt.Println("New loop")
		select {
		case <-a:
			aclosed = true // HL
		case <-b:
			bclosed = true
		}
	}
}

// STOP2 OMIT
