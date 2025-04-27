package main

import (
	"io"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Error:", r)
		}
	}()
	filename := "sample.txt"
	// creating a file with os package
	fileptr, err := os.Create(filename)
	ChecknilError(err)
	defer fileptr.Close()
	// writing to a file with io package
	legth, err := io.WriteString(fileptr, "Hello World")
	ChecknilError(err)
	println("Length of file is", legth)
	Readfile(filename)
}

func Readfile(filename string) {
	// reading a file with os package
	databypes, err := os.ReadFile(filename)
	ChecknilError(err)
	println("Data in file is", string(databypes))

}

func ChecknilError(err error) {
	if err != nil {
		panic(err)
	}
}
