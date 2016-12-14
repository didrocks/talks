// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// START OMIT
	const n = 100               // HL
	quit := make(chan struct{}) // HL
	var done sync.WaitGroup
	for i := 0; i < n; i++ {
		done.Add(1)
		go func() {
			defer done.Done()
			select {
			case <-time.After(1 * time.Hour):
			case <-quit:
			}
		}()
	}
	t0 := time.Now()
	close(quit) // HL
	done.Wait() // wait for the goroutine to stop
	fmt.Printf("Waited %v for %d goroutines to stop\n", time.Since(t0), n)
	// STOP OMIT
}
