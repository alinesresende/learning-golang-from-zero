package main

import (
	"fmt"
	"time"
)

// Pattern Generator

func main() {
	channel := write("Hi, world!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintln("Valor recebido: ", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
