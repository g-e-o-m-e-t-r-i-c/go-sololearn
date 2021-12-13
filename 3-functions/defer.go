package main

import "fmt"

func welcome() {
	fmt.Print("hi!")
}

func main() {
	// defer welcome() // will only be run after the main function has finished
	// fmt.Println("hello")

	fmt.Println("starting loop...")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // multiple defers will be stacked on top of each other, in LIFO order
	}
}
