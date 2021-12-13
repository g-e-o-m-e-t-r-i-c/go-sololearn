package main

import (
	"fmt"
	"time"
)

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result // send the result via channels
}

func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	ch <- result // send the result via channels
}

func main() {
	evench, sqch := make(chan int), make(chan int)

	go evenSum(0, 100, evench)
	go squareSum(0, 100, sqch)

	// fmt.Println(<-evench + <-sqch) // combines data from both channels

	// using select to select data from channels
	for { // will wait until one channel gets data. till then, will print "nothing available"
		select { // waits for a channel to receive data and executes its respective case
		// only one of the cases will be executed
		case x := <-evench:
			fmt.Println(x)
			return
		case y := <-sqch:
			fmt.Println(y)
			return
		default: // when no data is available, prevent deadlocks
			fmt.Println("nothing available")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
