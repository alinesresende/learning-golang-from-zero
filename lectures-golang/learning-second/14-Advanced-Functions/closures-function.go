package main

import "fmt"

// A closure function is a function that takes another function.
// The function "exampleFunction" is using a variable "text" that is outside the function.
func closure() func() {
	text := "Inside the closures function"

	exampleFunction := func() {
		fmt.Println(text)
	}
	return exampleFunction
}

func main() {
	// Closures Function
	// These are functions that reference variables that are outside their scope.
	text := "Inside the main function"
	fmt.Println(text)

	functionNew := closure()
	functionNew()
}
