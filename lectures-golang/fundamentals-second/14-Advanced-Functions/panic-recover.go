package main

import "fmt"

// Panic Function
// The Panic function interrupts the normal flow of the program and will stop everything
// Panic is different from error.
// In the error it is possible to continue the execution.
// In Panic it is not possible to continue the execution.

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution successfully recovered")
	}
}

func approvedStudent(n1, n2 float64) bool {
	// Defer in recover function for execution to continue
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
	// Possibility of panic execution
	panic("The average is exactly six")
}

func main() {
	fmt.Println(approvedStudent(6, 6))
	fmt.Println("post execution")
}
