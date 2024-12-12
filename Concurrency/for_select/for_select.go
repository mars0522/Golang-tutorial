package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	fmt.Println("hello world!")
}

func dosomework(done <-chan bool) {
	for {
		select {
		case <-done:
			// return
			go work()
			time.Sleep(time.Second * 10)
		default:
			fmt.Println("Hello world!")
		}
	}
}
func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	wg.Add(1)
	go dosomework(done)

	for {
		done <- true
	}

	// time.Sleep(time.Second * 3)
	// close(done)
	// wg.Wait()

}
