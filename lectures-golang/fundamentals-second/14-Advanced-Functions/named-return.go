package main

import "fmt"

// Functions with named return
// Return values ​​in Golang can be named and act like variables.
// The return statement with no arguments returns the current values ​​of the results. This is known as a "clean" return.

func mathematicalCalculations(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return // It is not necessary to write the name of the variables
}

func main() {
	sum, subtraction := mathematicalCalculations(10, 20)
	fmt.Println(sum, subtraction)
}
