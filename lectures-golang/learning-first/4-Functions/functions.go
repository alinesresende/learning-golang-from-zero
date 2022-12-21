package main

import "fmt"

// After the parameters specifies the return of the function
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Functions with more than one return
func calculations(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// Declare variable of type function and cast a function inside
	var exampleFunc = func() {
		fmt.Println("Function inside variable")
	}
	exampleFunc()

	// Another example

	var newExamplo = func(text string) {
		fmt.Println(text)
	}
	newExamplo("New Function")

	// when I don't want to use a variable I use "_"
	_, resultSubtraction := calculations(10, 15)
	fmt.Println(resultSubtraction)

	resultSum, resultSubtraction := calculations(10, 15)
	fmt.Println(resultSum, resultSubtraction)
}
