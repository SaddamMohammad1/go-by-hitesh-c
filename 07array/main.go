package main

import "fmt"

func main() {

	/*
		Array: Array in Go is a fixed-size collection of elements of the same data type. The size of an array is defined at declaration time and cannot be changed later.
	*/
	fmt.Println("Welcome to array in go lang")

	// Declaring an array of size 4 which can store string values
	// All elements are initialized with default value "" (empty string)
	var fruitList [4]string

	// Assigning values to specific indexes
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"

	// Index 2 is NOT assigned any value
	// So it will contain default value "" (empty string)
	fruitList[3] = "Peatch"

	fmt.Println("Fruit list is:", fruitList) // Output - Fruit list is: [Apple Tomato  Peatch]	Here after Tomator comes a space because here directly insert value in 1 to 3, here not assing any value to index 2 so bydefault for string empty string take

	fmt.Println("FruitList length is:", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegy list is:", vegList)

}
