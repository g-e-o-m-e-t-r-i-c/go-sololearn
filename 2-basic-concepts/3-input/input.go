package main

import "fmt"

func main() {
	var input string
	// takes in user input as the variable "input", then prints it out
	fmt.Scanln(&input) // note the & (returns the address of the variable)
	fmt.Println(input)
	// NOTE: the above Println must take place before the next Scanln,
	// thus the second input will not be scanned unti the first input is taken in and printed.

	var num int
	fmt.Scanln(&num) // takes in the input as an integer
	fmt.Println(num * 50)
}
