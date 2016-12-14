// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// START OMIT
	quit := make(chan bool)
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		defer done.Done()
		select {
		case <-time.After(1 * time.Hour):
		case <-quit:
		}
	}()
	t0 := time.Now()
	quit <- true // HL
	done.Wait()  // wait for the goroutine to stop
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
	// STOP OMIT
}
