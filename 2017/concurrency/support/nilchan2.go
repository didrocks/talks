// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// START OMIT
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // HL
	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() { b <- "b" }()
	if r.Intn(2) == 0 { // HL
		a = nil              // HL
		fmt.Println("nil a") // HL
	} else {
		b = nil
		fmt.Println("nil b")
	}
	select {
	case s := <-a:
		fmt.Println("got", s)
	case s := <-b:
		fmt.Println("got", s)
	}
	// STOP OMIT
}
