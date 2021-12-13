package main

import "fmt"

func main() {
	x, y := 42, 8 // set variables to be used


	/***********
	ARITHMETIC
	***********/

	// addition
	sum := x + y
	fmt.Println("Sum:", sum)

	// subtraction
	diff := x - y
	fmt.Println("Difference:", diff)

	// multiplication
	prod := x * y
	fmt.Println("Product:", prod)

	// integer division (returns only the quotient)
	quo := x / y
	fmt.Println("Quotient:", quo)

	// modulo (remainder after division)
	mod := x % y
	fmt.Println("Modulo:", mod)

	// short assignment
	c := 9
	d := 15
	c += d // new value of c: c + d, d remains
	fmt.Println("Short Assignment (Addition):", c)

	d -= c // new value of d: d - c, c remains
	fmt.Println("Short Assignment (Subtraction):", d)

	d *= c // new value of d: d * c, c remains
	fmt.Println("Short Assignment (Multiplication):", d)

	c /= d // new value of c: c / d, d remains
	fmt.Println("Short Assignment (Int. Divison):", c)

	d %= 6 // new value of d: d % 6, c remains
	fmt.Println("Short Assignment (Modulo):", d)

	// concatenation (combining two strings)
	a := "hi "
	b := "golang!"
	fmt.Println("Concatenation:", a + b)

	/*******************
	LOGICAL & RELATIONAL
	*******************/
	// logical and relational operators always return bool.
	e := 2
	f := 5
	fmt.Println("Equal:", e == f) // equal to
	fmt.Println("Not Equal:", e != f) // not equal to
	fmt.Println("Greater Than:", e > f) // greater than
	fmt.Println("Less Than:", e < f) // less than
	fmt.Println("Greater Than or Equal To:", e >= f) // greater than or equal to
	fmt.Println("Less Than or Equal To:", e <= f) // less than or equal to

	fmt.Println("Logical And:", e == f && e < f) // only true when both conditions are true
	fmt.Println("Logical Or:", e != f || e > f) // true when one of the two conditions is true, false when none of the conditions is true
	fmt.Println("Logical Not:", !(e >= f)) // reverses the outcome
}
