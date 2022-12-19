package main

import (
	"fmt"
)

func main() {
	// Arithmetic operators
	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	multiplication := 10 * 5
	restOfDivision := 10 % 2

	fmt.Println(sum, subtraction, division, multiplication, restOfDivision)
	// It is not possible to perform calculations with different types
	// example: sum type int8 + int32
	var number1 int16 = 10
	var number2 int16 = 25
	sum2 := number1 + number2
	fmt.Println(sum2)

	// Assignment Operators
	var text string = "String"
	text2 := "String2"
	fmt.Println(text, text2)

	// Relational Operators
	// return a boolean (true or false)
	fmt.Println(1 > 2)  // false
	fmt.Println(1 >= 2) // false
	fmt.Println(1 == 2) // false
	fmt.Println(1 <= 2) // true
	fmt.Println(1 < 2)  // true
	fmt.Println(1 != 2) // true

	// Logical Operators
	truth, lie := true, false
	fmt.Println(truth && lie) // false
	fmt.Println(truth || lie) // true
	fmt.Println(!truth)       // false
	fmt.Println(!lie)         // true

	// Unary operators
	numberOne := 10
	numberOne++
	numberOne += 15
	fmt.Println(numberOne)

	// there is no ternary operator in Golang
}
