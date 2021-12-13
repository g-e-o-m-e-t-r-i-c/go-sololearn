package main

import "fmt"

func main() {
	// declare mutable variables with "var"
	var i int = 8 // keyword, name, type
	var j, k int = 9, 57 // declare multiple variables of the same type
	var m = "hello" // type inference

	// short initialisation
	// name := value (:= is the assignment operator)
	n := 100 // declares and initialises variable n to the value 100

	// constant values (cannot be changed throughout the program)
	const pi float64 = 3.141592653589793 // cannot use the := operator (assignment) for consts

	fmt.Printf("%d %d %d %s %d\n", i, j, k, m, n);
}
