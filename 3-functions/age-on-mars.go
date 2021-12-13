package main

import "fmt"

func mars_age(earth_age int) int {
	return earth_age * 365 / 687
}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
