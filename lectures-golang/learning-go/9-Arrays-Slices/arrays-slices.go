package main

import "fmt"

func main() {
	// Array is a list of values
	// Ways to create arrays
	// All data within the array must be of the same type
	// The number inside the square brackets refers to the size of the list.
	// Array specification is mandatory
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "first position"
	fmt.Println(array2)

	array3 := [5]string{"first position", "second position", "third position", "fourth position", "fifth position"}
	fmt.Println(array3)
}
