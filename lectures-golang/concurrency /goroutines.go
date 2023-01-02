package main

import (
	"fmt"
	"time"
)

func main() {
	// competition != parallelism
	// Concurrency is basically the ability to handle multiple things at once while parallelism is the ability to handle multiple things at the same time.

	write("Hi, world!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
