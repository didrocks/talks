// +build OMIT

package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for i := 0; i < 10; i++ { // HL
		v, ok := <-ch
		fmt.Println(v, ok)
	}
	// STOP OMIT
}
