package main

import "fmt"

// Worker Pools
// A worker pool is basically a collection of threads that sit around waiting for tasks to be assigned to them.

func main() {
	tasks := make(chan int, 45) // channel with buffer
	result := make(chan int, 45)

	go worker(tasks, result)
	go worker(tasks, result)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-result
		fmt.Println(result)
	}
}

// Task channel only receives data -  The arrow "<-" comes before the word chan
// The result channel only sends data - The arrow "<-" comes after the word chan

func worker(tasks <-chan int, result chan<- int) {
	for number := range tasks {
		result <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}
