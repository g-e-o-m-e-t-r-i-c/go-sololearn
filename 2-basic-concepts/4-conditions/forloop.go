package main

import "fmt"

func main() {
	// init, cond, post
	for i := 0; i < 5; i++ {
		fmt.Println(i);
	}

	// init and post can be omitted (similar to while loop)
	fac := 1
	counter := 1
	for counter <= 15 {
		fac *= counter
		counter++
	}

	fmt.Printf("15! = %d", fac)
}
