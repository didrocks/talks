// +build OMIT

package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan int, 2)
	close(ch)
	ch <- 3 // HL

	for v := range ch {
		fmt.Println(v)
	}
	// STOP OMIT
}
