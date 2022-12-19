package main

import (
	"errors"
	"fmt"
)

func main() {
	// four kind of integers
	// int8, int16, int32, int64
	// int alone uses your architecture
	// int also support negative numbers
	var numbers int16 = 100
	fmt.Println(numbers)

	var number1 int64 = -100000
	fmt.Println(number1)

	// uint means unsigned, unsigned
	// int36 = rune
	// uint8 = byte
	var number2 uint32 = 1500
	fmt.Println(number2)

	// real numbers
	// float32, float64
	// support numbers broken with comma and period
	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	realNumber2 := 12345.67
	fmt.Println(realNumber2)

	//types string
	var str string = "text"
	fmt.Println(str)

	str2 := "text2"
	fmt.Println(str2)

	// Zero Value
	// Value assigned to uninitialized variable
	// Return is a blank string
	var text string
	fmt.Println(text)

	// Return of a blank number is zero
	var numbers1 int16
	fmt.Println(numbers1)

	// Types booleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	//absence of boolean declaration results in false
	var booleano2 bool
	fmt.Println(booleano2)

	// Types Null
	// Type is error
	//absence of error declaration is nil
	var erro error
	fmt.Println(erro)

	// Sintexa do error
	var erro1 error = errors.New("Text incorrect")
	fmt.Println(erro1)
}
