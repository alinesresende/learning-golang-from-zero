package main

import (
	"fmt"
	"time"
)

func main() {
	// Repeating structure
	// First example of FOR
	number := 0
	for number < 10 {
		number++
		fmt.Println("Incrementing number")
		time.Sleep(time.Second)
	}

	// Secund example of FOR
	for numberTwo := 0; numberTwo < 10; numberTwo++ {
		fmt.Println("Incrementing numberTwo", numberTwo)
		time.Sleep(time.Second)
	}
}
