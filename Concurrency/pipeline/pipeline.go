package main

import "fmt"

func sliceToChannel(arr []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range arr {
			ch <- val
		}
		close(ch)
	}()
	return ch
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}
func main() {

	numbers := []int{1, 2, 3, 4, 5}
	channel1 := sliceToChannel(numbers)
	channel2 := sq(channel1)

	for val := range channel2 {
		fmt.Println(val)
	}

}
