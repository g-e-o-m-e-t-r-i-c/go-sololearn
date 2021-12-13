package main

import "fmt"

func main() {
	var a int = 90 // integer
	var b float32 = 3.14 // float with single-precision
	var c float64 = 3.14159265359 // float with double-precision
	var d string = "hello" // string
	var e bool = true // boolean

	// default values - when you initialise a variable without its value
	var f int // default: 0
	var g string // default: ""
	var h bool // default: false

	fmt.Println("Integer: ", a)
	fmt.Println("Float (Single-precision): ", b)
	fmt.Println("Float (Double-precision): ", c)
	fmt.Println("String: ", d)
	fmt.Println("Boolean: ", e)
}
