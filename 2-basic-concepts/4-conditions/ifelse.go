package main

import "fmt"

func main() {
	var age int // declares variable x
	fmt.Print("enter your age: ") // not Println, we don't want a newline after the prompt
	fmt.Scanln(&age) // input age

	if age >= 18 { // if statement
		fmt.Println("we have nice booze here")
	} else if age == 17 { // else if (if above condition not satisfied)
		fmt.Println("wait one more year yp")
	} else { // else (if all above conditions not satisfied)
		fmt.Println("wut are you doin here kiddo")
	}
}
