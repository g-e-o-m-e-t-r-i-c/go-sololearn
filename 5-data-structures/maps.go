package main

import "fmt"

func main() {
	// declare a map using make() method
	m := make(map[string]int) // key is a string, value is an integer

	// mapping strings to int
	m["james"] = 26
	m["bob"] = 34
	m["mary"] = 28

	// alternative syntax
	n := map[string]int {
		"james": 26,
		"bob": 34,
		"mary": 28, // leave a trailing comma
	}

	// output values using keys
	fmt.Println("mary's age:", n["mary"])

	// remove key: value pair from map
	delete(m, "bob") // removes key "bob" and value 34 from map m

	// reassigning values
	m["mary"] = 25

}
