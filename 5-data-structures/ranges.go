// range-based for loop
package main

import "fmt"

func main() {
	a := [5]int {5, 7, 4, 6, 9} // declare an array of fixed size

	// every time the loop runs, it returns the index and the value at that index
	// thus, two variables are required
	fmt.Println("indices with values")
	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	// if the index is not required, an underscore can be used as its variable
	fmt.Println("values only")
	for _, value := range a {
		fmt.Printf("%d ", value)
	}

	// ranges can also be used for strings
	// a. unicode codepoints
	var str string = "Hello World"
	for _, char := range str {
		fmt.Println(char) // will print a series of integers denoting unicode codepoints for each character
	}

	// b. characters in the string
	for _, char := range str {
		fmt.Printf("%c ", char)
	}
}
