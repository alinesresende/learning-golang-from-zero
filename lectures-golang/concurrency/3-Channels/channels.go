package main

import (
	"fmt"
	"time"
)

// Channel
// The channel can send and receive data to organize the goroutines
// To create a channel it is necessary to use the make() function, with the word "chan" and then specify a type (ex: string)
// The use of "<-" serves to send something

func main() {
	channel := make(chan string)
	go write("hi, world!", channel)

	fmt.Println("After the function starts executing")

	for message := range channel { // receiving channel
		fmt.Println(message)
	}

	fmt.Println("End of program")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // sending to channel
		time.Sleep(time.Second)
	}
	close(channel)
}
