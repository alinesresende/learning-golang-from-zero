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
	slice = append(slice, 6)
	fmt.Println(len(slice)) // length
	fmt.Println(cap(slice)) // capacity
	fmt.Println(slice)
	fmt.Println("--------------")

	slice1 := make([]float32, 3)
	fmt.Println(len(slice1)) // length
	fmt.Println(cap(slice1)) // capacity
	fmt.Println("--------------")

	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)
	slice1 = append(slice1, 6)
	slice1 = append(slice1, 7)
	fmt.Println(len(slice1)) // length
	fmt.Println(cap(slice1)) // capacity
	fmt.Println(slice1)

}
