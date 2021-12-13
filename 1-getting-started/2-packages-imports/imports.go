// Main program to be run
package main

// Import multiple packages at once
import (
	"fmt"
	"math"
)

import "time" // Import one package by itself

func main() {
	// Main code goes here
	// Print statement
	fmt.Println("hi")

	// Math
	const a = 3 // set constants
	const b = 4
	c := math.Sqrt(a*a + b*b) // assigns c to the value of sqrt(a^2 + b^2)
	fmt.Printf("%.1f\n", c)   // prints output to 1 d.p.

	// Time
	start := time.Now()
	fmt.Println("Started time!")
	end := time.Now()
	fmt.Println("Stopped Time!")

	elapsed := end.Sub(start)
	fmt.Println("Started: ", start, "\nEnded: ", end, "\nElapsed: ", elapsed)
}
