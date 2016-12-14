// +build OMIT

package main

func main() {
	// START1 OMIT
	var ch chan bool
	ch <- true // blocks forever
	<-ch       // blocks forever as well
	// STOP1 OMIT
}
