package main

import "fmt"

func change_ptr(ptr *int) { // declare function using pointer
	*ptr = 9
}

func main() {
	var x int = 42 // declares integer x
	var p *int = &x // declares pointer p, which references x's address
	fmt.Println("Value of x:", x) // prints value
	fmt.Println("Address of x:", p) // prints address
	fmt.Println("Value which the address is referencing:",*p) // prints value referenced by address

	// change the value of x using pointers
	*p = 8
	fmt.Println("New value of x:", x)

	change_ptr(p)
	fmt.Println("changed value of pointer address:", *p)
}
