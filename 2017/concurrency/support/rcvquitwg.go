// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func cleanup() {
	fmt.Println("Cleanup function")
}

// START1 OMIT
func main() {
	done := &sync.WaitGroup{} // HL
	quit := make(chan string)
	done.Add(1)                    // HL
	c := boring("Joe", quit, done) // HL
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	done.Wait() // HL
	fmt.Println("Everything ended properly")
}

// STOP1 OMIT

func boring(msg string, quit <-chan string, done *sync.WaitGroup) <-chan string {
	c := make(chan string) // HL
	// START2 OMIT
	go func() {
		defer done.Done() // HL
		defer cleanup()   // HL
		// STOP2 OMIT
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
				// START3 OMIT
			case <-quit: // HL
				fmt.Println("Bye bye")
				return
				// STOP3 OMIT
			}
		}
	}()
	return c
}
