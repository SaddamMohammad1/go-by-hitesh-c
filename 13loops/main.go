package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// Here return index
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// Here return index and value also
	for index, value := range days {
		fmt.Printf("Index is %v and value is %v \n", index, value)
	}

	// It's similar to While loop
	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 5 {
			break
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}
}
