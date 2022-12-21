package main

import "fmt"

func main() {
	// Evaluation of a condition
	// Basic if/else Structure
	number := -5

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less than 15")
	}

	// Example if init
	// If with a variable inside
	// In this example the anoterNumber variable is receiving the number variable
	// Cannot access anotherNumber variable outside the If structure
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Number greater than 0")
	} else {
		fmt.Println(("Number less than 0"))
	}
}
