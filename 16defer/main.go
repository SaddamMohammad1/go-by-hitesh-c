package main

import "fmt"

func main() {

	// defer = executes at the end of the function (main)
	// Multiple defer follow LIFO (Last In First Out)

	defer fmt.Println("World") // Executes last
	defer fmt.Println("One")   // Executes before World
	defer fmt.Println("Two")   // Executes first among these

	fmt.Println("Hello") // Executes immediately

	myDefer() // Calls function having defer inside loop
}

func myDefer() {

	// defer inside loop stores each value of i
	// Execution happens after function ends
	// Output will be in reverse order (stack behavior)

	for i := 0; i < 5; i++ {
		defer fmt.Print(i) // Stores 0,1,2,3,4 → prints 4,3,2,1,0
	}
}
