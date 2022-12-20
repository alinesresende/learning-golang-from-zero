package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array is a list of values
	// Ways to create arrays
	// All data within the array must be of the same type
	// The number inside the "[ ]" refers to the size of the list.
	// Array specification by "[ ]" is mandatory
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "first position"
	fmt.Println(array2)

	array3 := [5]string{"first position", "second position", "third position", "fourth position", "fifth position"}
	fmt.Println(array3)

	// when not specifying size you can use "..."
	// However, using "..." does not make the array dynamic.
	array4 := [...]int{1, 2, 3, 4, 5, 60}
	fmt.Println(array4)

	// Slice
	// Slice is an array-based structure with some flexibility
	// In Slice it is not necessary to specify the size inside the "[]"
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array3)) //[5]int

	// In Slice there is a function called append to return a new Slice
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = "changed position"
	fmt.Println(slice2)
}
