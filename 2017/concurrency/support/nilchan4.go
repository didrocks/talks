// +build OMIT

package main

import "fmt"
import "time"

func main() {
	// START1 OMIT
	a, b := make(chan bool), make(chan bool)
	go func() {
		close(a)
		time.Sleep(time.Second) // HL
		close(b)
	}()
	WaitMany(a, b)
	// STOP1 OMIT
}

// START2 OMIT
// WaitMany waits for a and b to close.
func WaitMany(a, b chan bool) {
	for a != nil || b != nil { // HL
		fmt.Println("New loop")
		select {
		case <-a:
			a = nil // HL
		case <-b:
			b = nil // HL
		}
	}
}

// STOP2 OMIT
