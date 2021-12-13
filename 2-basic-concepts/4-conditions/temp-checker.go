package main

import "fmt"

func main() {
	var temp float32
	fmt.Scanln(&temp)

	if temp > 99.5 {
		fmt.Println("Fever")
	} else {
		fmt.Println("Allowed")
	}
}
