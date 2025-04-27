package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)       // make a reader to read from standard input
	fmt.Print("Enter text: ")                 // prompt user for input
	userInput, err := reader.ReadString('\n') // read a line from standard input till new line
	if err != nil {
		panic(err)
	}
	fmt.Println("You entered:", userInput)
}
