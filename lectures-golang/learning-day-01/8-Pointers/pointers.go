package main

import "fmt"

func main() {
	// Pointer is a memory reference
	// Use of the "&"" symbol means that the pointer variable is storing the variable address
	// The value received is the address of the variable
	var number int
	var pointer *int
	number = 100
	pointer = &number
	fmt.Println(number, *pointer)

	// Another example

	var creature string = "shark"
	var pointer1 *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer1)

	fmt.Println("*pointer =", *pointer1)

	*pointer1 = "jellyfish"
	fmt.Println("*pointer =", pointer1)
	fmt.Println("creature =", creature)
}
