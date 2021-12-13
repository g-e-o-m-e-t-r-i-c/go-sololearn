package main

import "fmt"

func main() {
	num := 3
	switch num { // unlike c, break statement is not needed
	case 1:
		fmt.Println("# is 1")
	case 2:
		fmt.Println("# is 2")
	case 3:
		fmt.Println("# is 3")
	default: // default case is optional
		fmt.Println("idk this number") // if above conditions not satisfied
	}

	// NOTE a switch using conditions (i.e. x < 20, x > y || x < y etc.)
	// does not take a parameter
	// hence the syntax is
	// switch { ... }
}
