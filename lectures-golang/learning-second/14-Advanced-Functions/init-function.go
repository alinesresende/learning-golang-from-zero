package main

import "fmt"

// Init Function
// The init() function causes a piece of code to run before any other part of your package.
// The init() function is executed before calling the main() function.

var n int

func init() {
	fmt.Println("Executing the init function")
	n = 10
}

func main() {
	fmt.Println("Executing the main function")
	fmt.Println(n)
}
