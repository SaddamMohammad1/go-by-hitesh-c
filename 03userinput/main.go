package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// ------------------------------------------------------
	// INPUT HANDLING IN GO (Using bufio.Reader)
	// ------------------------------------------------------

	// Create a new buffered reader to read input from standard input (Stdin)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// This is called "comma ok / comma error syntax" in Go
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

	// NOTE:
	// - Input read from console is always in STRING format
	// - Newline character '\n' is included in the input
	// - Conversion is required if numeric value is needed (strconv package)
}
