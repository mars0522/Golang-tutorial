package main

import (
	"bufio"
	"concurrency/Basics/project1/note"
	"fmt"
	"os"
	"strings"
)

func main() {

	title := getUserInput("Enter title:")
	content := getUserInput("Enter content:")
	new_note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Save(new_note)

}

func getUserInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "")
	return input
}
