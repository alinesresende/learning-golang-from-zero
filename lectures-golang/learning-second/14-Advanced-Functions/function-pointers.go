package main

import "fmt"

func invertSign(number int) int {
	return number * -1
}

func invertSignWithPointers(number *int) {
	*number = *number * -1
}

func main() {
	// Function with pointers
	number := 20
	invertNumber := invertSign(number)
	fmt.Println(invertNumber)

	fmt.Println("-----------")

	newNumber := 40
	fmt.Println(newNumber)
	invertSignWithPointers((&newNumber))
	fmt.Println(newNumber)
}
