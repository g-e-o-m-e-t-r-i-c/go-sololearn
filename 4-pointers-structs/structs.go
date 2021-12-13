package main

import "fmt"

// syntax:
/*
	type [name] struct {
		field datatype
		field datatype
		...
	}
*/

type Contact struct { // define a new datatype of Contact
	name string
	age int
	phone string
}

// a method tied to a struct
/*
	func (param struct) func_name() {
		// stuff
	}
*/
func (x Contact) info() { // the struct is the receiver
	fmt.Printf("hi, %s, it says here that you're %d this year, and your number is %s. is that right?\n", x.name, x.age, x.phone)
}

func (y Contact) get_hp() string { // add a return type to a method
	return y.phone
}

// pointers in method receivers
// func (ptr_name *struct) method_name(param) return_type {
// }
// NOTE: only pointers can be used if you want to change the value of any struct
func (ptr *Contact) change_age(change int) int {
	return ptr.age + change // NOTE that struct ptrs can auto-deference
}

func main() {
	// define a new instance of a struct
	// syntax: [var_name] := [struct_name] { field1, field2, field... }
	james := Contact {"James", 29, "98605784"} // define a new Contact, with a name, age and hp. no.

	// accessing the data in a struct using dot syntax
	// syntax: struct.field
	fmt.Println("name:", james.name)
	fmt.Println("age:", james.age)
	fmt.Println("phone number:", james.phone)

	// pointers to structs
	emily := Contact {"Emily", 26, "95403265"}
	em := &emily // define a new pointer referencing "emily" contact

	// pointers to structs are automatically deferenced, so this syntax is valid
	// the following won't print a memory address, but rather a value
	fmt.Println("emily's age:", em.age) // using the pointer in place of the actual struct instance

	// pointers can be made automatically while creating a struct
	bryan := &Contact {"Bryan", 32, "98824510"}
	fmt.Println("bryan's hp:", bryan.phone) // will print the actual hp no., not a memory address

	// use dot syntax to access methods
	james.info() // accesses info method of struct instance james
	fmt.Println("james's hp:", james.get_hp()) // using return types in methods

	// we can still use the original struct instance "james" instead of using a pointer
	fmt.Println("james's new age:", james.change_age(5)) // adds 5 to james's age
}
