package main

import "fmt"

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "one"
	}()

	go func() {
		chan2 <- "two"
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println("received message from chan1: ", msg1)
	case msg2 := <-chan2:
		fmt.Println("received message from chan2: ", msg2)
	}

}
