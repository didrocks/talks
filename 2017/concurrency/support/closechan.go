// +build OMIT

package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch) // HL

	for v := range ch {
		fmt.Println(v) // called twice
	}
	// STOP OMIT
}
