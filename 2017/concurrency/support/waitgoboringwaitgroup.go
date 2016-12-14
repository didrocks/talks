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
	var done sync.WaitGroup     // HL
	done.Add(1)                 // HL
	go boring("boring!", &done) // HL
	fmt.Println("I'm listening.")

	done.Wait() // HL
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, done *sync.WaitGroup) { //HL
	defer done.Done() // HL
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		if i == 10 {
			fmt.Println("I'm borely done")
			return
		}
	}
}

// STOP OMIT
