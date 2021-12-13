package main

import "fmt"

func main() {
	var f int
	fmt.Scanln(&f)
	switch {
	case f < 0:
		fmt.Println("Wrong Input")
	case f < 20:
		fmt.Println("Infrasound")
	case f >= 20 && f <= 20000:
		fmt.Println("Audible")
	case f > 20000:
		fmt.Println("Ultrasound")
	}
}
