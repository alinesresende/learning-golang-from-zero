package main

import "fmt"

// A recursive function is when one or more calls to itself are made within it.

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	position := uint(5)

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}
