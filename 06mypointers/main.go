package main

import "fmt"

func main() {
	// Pointer is a variable that stores the memory address of another variable instead of storing the actual value.
	// Pointers help in memory efficiency, modifying values directly, and passing large data without copying.
	fmt.Println("Welcome to a class on pointers")

	// Declaring a pointer variable of type *int
	// At this point, it is not pointing to any memory location
	var ptr *int
	fmt.Println("Value of pointer is:", ptr) // Output - <nil>

	myNumber := 23
	// &myNumber gives the memory address of myNumber
	// newPtr stores the address of myNumber
	var newPtr = &myNumber

	// Printing the memory address stored in newPtr
	fmt.Println("Value of pointer is:", newPtr) // Output - 0xc00000a098

	// Dereferencing the pointer using *
	// This gives the actual value stored at that memory address
	fmt.Println("Value of pointer is:", *newPtr) // Output - 23

}
