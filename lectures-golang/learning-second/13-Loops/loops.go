package main

import (
	"fmt"
)

func main() {
	// Repeating structure
	// First example of FOR

	number := 0
	for number < 10 {
		number++
		fmt.Println("Incrementing number")
	}

	// // Secund example of FOR
	// // The increment can be with other numbers

	for numberTwo := 0; numberTwo < 10; numberTwo++ {
		fmt.Println("Incrementing numberTwo", numberTwo)
	}

	// FOR with range clause
	// The range clause is used to iterate over something
	// Example: map, string...
	// The index and value declaration is mandatory
	// If you don't want to use the index or value, you need to put a "_" in its place

	names := [3]string{"Aline", "Thais", "Vera"}
	for index, name := range names {
		fmt.Println(index, name)
	}

	// Interaction by String
	// it is necessary to put String to be able to visualize the content

	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	// Interaction by Map

	user := map[string]string{
		"name":      "Aline",
		"lastename": "Resende",
	}
	for key, value := range user {
		fmt.Println(key, value)
	}

	// note: Unable to range in Struct
}
