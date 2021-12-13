package main

import "fmt"

func main() {
	// declaring an array
	var a [5]int // declares an array arr of size 5

	// short declaration using assignment operator
	b := [5]int{0, 2, 4, 6, 8}

	// array indices
	for i := 0; i < 5; i++ {
		a[i] = i*2 + 1
	}

	fmt.Println("4th element of b:", b[3])

	// array slicing (similar to python)
	// array slicing returns a range inside the array, including the lower bound, but not the higher bound.
	fmt.Println("1st - 3rd elements of a:", a[0:3]) // will print out a[0], a[1] and a[2] but NOT a[3]

	// can omit the lower bound if it is the first element,
	// can omit the higher bound if it is the last element
	fmt.Println("first 2 elements of b:", b[:2])
	fmt.Println("last 3 elements of a:", a[2:])

	// since slices return a smaller array, we can use the same indexing system for slices
	s := a[1:4]                               // will get middle 3 elements of a
	fmt.Println("second element of s:", s[1]) // aka the 3rd element of a

	// alternative syntax
	fmt.Println("third element of a (using slices):", a[1:4][1])

	// make() function to create slices
	c := make([]int, 5) // dynamically sized array with size 5 at first
	// NOTE: original size can be set to 0, so that appending will not result in any leading zeros
	for i := 1; i <= 6; i++ {
		c = append(c, i) // append items to dynamically sized array
	}
	fmt.Println("new dynamic array:", c) // first 5 elements will be filled with 0s, as those are the first 5 originally created, anything else will just be appended
}
