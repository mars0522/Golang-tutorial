package main

import (
	"fmt"
)

func main() {
	// channels are used to facilitate the communication between the goroutines
	// Channels used FIFO Data structure

	mychal := make(chan string)

	go func() {
		mychal <- "Hello World!"
	}()

	msg := <-mychal // this line of code is blocking
	// Main function is waiting either to receive the data from the channel or channel to be closed
	fmt.Println("Data from the channel: ", msg)
}
