package main

import (
	"fmt"
	"time"
)

// Select
// The select clause is used so that a function can work with multiple channels
// The select clause blocks the execution of the function until one of the channels is ready to be executed.
// If more than one channel is ready to run, it will randomly select which one to run.

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "channel one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "channel two"
		}
	}()

	for {
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}
}
