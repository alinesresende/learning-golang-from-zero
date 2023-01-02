package main

import (
	"fmt"
	"time"
)

func main() {
	// competition != parallelism
	// In parallelism, tasks are executed at the same time.
	// In concurrency, tasks may or may not be executed at the same time.

	write("Hi, world!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
