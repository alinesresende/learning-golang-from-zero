package main

import "fmt"

func funcFirst() {
	fmt.Println("Executing function 1")
}

func funcSecund() {
	fmt.Println("Executing function 2")
}

func approvedStudent(n1, n2 float32) bool {
	// The defer is performed before the return
	defer fmt.Println("Average calculated. Result will be returned")
	fmt.Println("Entering the function to check if the student has passed")
	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}
	return false
}

func main() {
	// Defer clause
	// The defer clause is responsible for deferring a piece of code
	// Using "defer" in FuncFirst defers function execution.
	// In functions that have a "defer" return, it is executed before
	defer funcFirst()
	funcSecund()
	fmt.Println(approvedStudent(3, 4))

}
