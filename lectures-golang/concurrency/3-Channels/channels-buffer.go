package main

import (
	"fmt"
)

// Channel with buffer
// syntax:
// ch := make(chan type, capacity)
// For the channel to be buffered, the capacity value must be greater than zero
// By default, channels are unbuffered, which means they will only accept sends (chan <-) if there is a matching receive (<- chan) that is ready to receive the sent value
// Buffered channel will only block execution when it reaches maximum capacity

func main() {
	channel := make(chan string, 2)

	channel <- "Hi, world!"            // sending to the channel
	channel <- "Programming in Golang" // sending to the channel

	message := <-channel // receiving the channel
	fmt.Println(message)
	fmt.Println(message)
}
