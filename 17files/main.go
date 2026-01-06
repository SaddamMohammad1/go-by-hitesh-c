package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// File handling in Go (Create → Write → Read)

	fmt.Println("Welcome to files in go lang")

	// Content to be written into file
	content := "This need to go in a file - LearnCodeOnline.in"

	// Create a new file (overwrites if already exists)
	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	// Write string data into file
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	// Shows number of bytes written
	fmt.Println("Length is:", length)

	// defer ensures file is closed at end of main()
	defer file.Close()

	// Read file content
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {

	// Read entire file into byte slice
	datatype, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	// Print raw byte data
	fmt.Println("Text data inside the file is \n", datatype)

	// Convert bytes to string for readable output
	fmt.Println("Text data inside the file is \n", string(datatype))
}

func checkNilErr(err error) {

	// Common error handler
	if err != nil {
		panic(err)
	}
}
