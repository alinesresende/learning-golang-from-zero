package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("hi, world!", channel)

	message := <-channel
	fmt.Println(message)
}

func write(text string, channel chan string) {
	for i := 0; 1 < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
}
