package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in go lang")
	grreter()

	// Call adder function
	result := adder(3, 5)
	fmt.Println("Result is:", result)

	// Here called proAdder function
	proResult := proAdder(3, 5, 6, 6, 7)
	fmt.Println("Pro Result is:", proResult)

}

// Funaction always define outside the main() function, In main() function only call.
func grreter() {
	fmt.Println("Welcome to go lang bro")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

// Here in function definition send any number of args like rest operator in JS
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
