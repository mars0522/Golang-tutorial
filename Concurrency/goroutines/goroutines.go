package main

import (
	"fmt"
	"time"
)

/*
* In the topic of concurrency in golang we will mostly encounter 3 Topics
* 1) Goroutines
* 2) Channels
* 3) Select
 */

func do1() {
	fmt.Println("1")
}
func do2() {
	fmt.Println("2")
}
func do3() {
	fmt.Println("3")
}
func main() {
	go do1()
	go do2()
	go do3()
	time.Sleep(time.Second * 2) // Point of join in the main function
	fmt.Println("Hello world!")
}
