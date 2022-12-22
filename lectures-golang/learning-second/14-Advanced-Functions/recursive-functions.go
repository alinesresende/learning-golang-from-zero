package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	position := uint(15)
	fmt.Println(fibonacci(position))

}
