package main

import (
	"fmt"
	"time"
)

func main() {
	// concurrency != parallelism
	// Concurrency is the ability to handle several things at once
	// Parallelism is the ability to handle multiple things at the same time
	// Goroutines are functions or methods that run concurrently
	// To use Goroutines it is necessary to include "go" before the function.

	go write("Hi, world!") // using goroutines
	write("Programming in golang")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
