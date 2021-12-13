package main

import "fmt"

func hi(name string) {
	fmt.Printf("hi %s", name)
}

/*
	func name(param type, param type) returnType {
		// stuff
	}
*/

func sum(a, b int) int {
	return a + b
}

// multiple return types
func swapStrInt(s string, n int) (int, string) {
	return n, s
}

func main() {
	var name string
	fmt.Scanln(&name)
	hi(name)

	var a, b int
	fmt.Print("enter a number: ")
	fmt.Scanln(&a)
	fmt.Print("enter another number: ")
	fmt.Scanln(&b)

	fmt.Printf("sum of 2 numbers: %d\n", sum(a, b))
}
