package main

import "fmt"

func main() {
	// Anonymous Functions
	// Anonymous function has no name
	// To execute the function it is necessary to use "()" after the end of the structure

	func() {
		fmt.Println("Olá Mundo!")
	}()

	// The anonymous function can take parameters

	func(text string) {
		fmt.Println(text)
	}("passing parameter")

	// The anonymous function can return
	// To print the return on the screen it is necessary to store the function inside a variable
	// Sprintf(): formats according to a format specifier and returns the resulting string.

	returnFist := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("passing parameter")

	fmt.Println(returnFist) // printing function return
}
