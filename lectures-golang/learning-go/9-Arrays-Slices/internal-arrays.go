package main

import "fmt"

func main() {
	// Internal Arrays
	// Syntax
	// function make: type + length + capacity
	// slice_name := make([]type, length, capacity)
	// The make function creates a slice with the informed capacity and returns the length
	// If the capacity parameter is not defined, it will be equal to length.
	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice)) // length
	fmt.Println(cap(slice)) // capacity

	slice = append(slice, 5)
	fmt.Println(slice)
}
