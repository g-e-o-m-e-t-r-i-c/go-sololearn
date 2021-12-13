// goroutines
package main

import (
	"fmt"
	"time"
)

// normal implementation (requires time.Sleep for concurrency)
func out(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond) // as an example
		fmt.Println(i)
	}
}

// better implementation (with channels)
func outch(from int, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond) // as an example
		fmt.Println(i)
	}
	ch <- true
}

func main() {
	// classic sequential program - the program moves one instruction by one
	// the first instruction needs to be completed before the second can even begin to be computed
	// out(0, 5)
	// out(6, 10)

	// goroutines (computation happens concurrently)
	// go doesn't need to wait for the first function to finish before starting the second
	// there will be no output in this scenario, as the main() function finishes before the out() calls
	fmt.Println("normal implementation without channel")
	go out(0, 5)
	go out(6, 10)

	// to fix the "no output" problem, we can add a time.Sleep function
	time.Sleep(500 * time.Millisecond)

	// the output is still jumbled up because multiple virtual threads are running at the same time, independent of each other

	// channels - allow goroutines to share data and communicate between each other
	// channel := make(chan int) // declare a channel

	// send data to the channel
	// channel <- 8 // sends 8 to the channel of ints

	// receive data from channel
	// n := <- channel // declaring a variable with the data from channel
	// fmt.Println(<- channel) // if we do not need a variable

	// better out() function with channels
	fmt.Println("better implementation with channel")
	ch := make(chan bool)
	go outch(0, 5, ch)
	go outch(6, 10, ch)
	<- ch // since main() is waiting to receive data from channel, it waits for outch() to finish before exiting


}
