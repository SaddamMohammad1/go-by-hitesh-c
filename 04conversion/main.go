package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	// ------------------------------------------------------
	// USER INPUT + STRING TO FLOAT CONVERSION IN GO
	// ------------------------------------------------------

	// Create a buffered reader to read input from the Stdin (standard input)
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// Convert string input to float64
	// strings.TrimSpace() removes leading/trailing spaces and newline characters
	// strconv.ParseFloat() converts string to float value
	// 64 specifies the precision (float64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
