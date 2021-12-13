package main

import "fmt"

// functions which take any number of arguments
func sum(nums ...int) int { // insert an ellipsis at the front to denote variable number of args
	total := 0
	for _, num := range nums { // use range to loop through the arguments
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(2, 5, 6))
	fmt.Println(sum(39, -45, 69, 184, 250))
	// ...

	// take a slice/array as an argument for a variadic function
	s := []int{3, 7, 6, 9, 10, 12}
	fmt.Println(sum(s...)) // append the ellipsis at the back
}
