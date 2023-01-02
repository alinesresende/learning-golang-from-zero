package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//WaitGroup is a very useful feature to synchronize all concurrent executions.
	var waitGroup sync.WaitGroup

	// adding two goroutines
	waitGroup.Add(2)

	// Usign anonymous function with goroutines
	go func() {
		write("Hi, world!")
		waitGroup.Done()
	}()

	go func() {
		write("Programming in Golang!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
