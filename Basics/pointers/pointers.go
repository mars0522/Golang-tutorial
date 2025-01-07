package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("You are an adult since : %d years\n", getAdultSince(&age))

}

func getAdultSince(age *int) int {
	return *age - 18
}
