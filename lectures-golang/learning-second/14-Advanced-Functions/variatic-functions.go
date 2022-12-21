package main

import "fmt"

// Variatic Functions
// They are those functions that have a variable number of parameters.
// When we use the symbol "..." we indicate that the number of parameters is variable.

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// Combination of function that receives fixed parameters with varied
// it is necessary to respect the order

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)

	}
}

func main() {
	resultSum := sum(1, 2, 3, 15, 27)
	fmt.Println(resultSum)
	write("Hello world", 10, 2, 3, 5, 6, 7)

}
