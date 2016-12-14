// +build OMIT
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	c := make(chan bool)
	go boring("boring1!", c)
    go boring("boring2!", c)
	fmt.Println("I'm listening.")

	<-c // HL
    <-c // HL
	fmt.Println("You're boring; I'm leaving.")
}
// STOP OMIT


func boring(msg string, c chan bool) {
	defer func() { c <- true }() // HL
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		if i == 10 {
			fmt.Println("I'm borely done")
			return
		}
	}
}

