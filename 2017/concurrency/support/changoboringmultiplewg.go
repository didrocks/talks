// +build OMIT
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func main() {
	var done sync.WaitGroup
	done.Add(2) // HL
	go boring("boring1!", &done)
	go boring("boring2!", &done)
	fmt.Println("I'm listening.")

	done.Wait()
	fmt.Println("You're boring; I'm leaving.")
}

// STOP OMIT

func boring(msg string, done *sync.WaitGroup) {
	defer done.Done()
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		if i == 10 {
			fmt.Println("I'm borely done")
			return
		}
	}
}
