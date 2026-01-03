package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) // output - [Apple Tomato Peach Mango Banana]

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) // output - [Tomato Peach Mango Banana]

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) // output - [Tomato Peach]

	highScores := make([]int, 4)

	highScores[0] = 1234
	highScores[1] = 5354
	highScores[2] = 3534
	highScores[3] = 454
	fmt.Println(highScores) // output - [1234 5354 3534 454]

	// Here declare 4 valur for inilization but we can add some more value using append method
	highScores = append(highScores, 3423, 243, 2432)
	fmt.Println(highScores) // output - [1234 5354 3534 454 3423 243 2432]

	// This is use for sort the slice values
	sort.Ints(highScores)
	fmt.Println(highScores) // output - [243 454 1234 2432 3423 3534 5354]

	/* How to remove a value from slices based on index */
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	// Note: append used in both condition for adding and removing value also
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // [reactjs javascript python ruby]
}
