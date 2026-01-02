package main

import "fmt"

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4553452453843
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit way of declaring a variable
	var website = "learncodeonline.in"
	fmt.Println(website)

	// Another way to declare a variable without var key
	numberOfUser := 300000 // That is called valarouse operator	(:=)
	fmt.Println(numberOfUser)

	/*
		valarouse operator (:=)
			This operator only use inside a function or block, not outside the function or block.
	*/
}
