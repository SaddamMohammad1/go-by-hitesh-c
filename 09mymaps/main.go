package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages) // List of all languages: map[JS:Javascript PY:Python RB:Ruby]

	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages) // List of all languages: map[JS:Javascript PY:Python]

	// Loops are interesting in go lang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// If you want to print only value then use _
	for _, value := range languages {
		fmt.Printf("For key, value is %v\n", value)
	}
}
