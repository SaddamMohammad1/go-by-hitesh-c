package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Small & safe URL for demo purpose
// example.com is reserved for documentation & testing
const url = "https://example.com"

func main() {

	// Simple HTTP GET request in Go
	fmt.Println("Making a web request...")

	// Send GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		panic(err) // Stop program if request fails
	}

	// response is of type *http.Response
	fmt.Printf("Response type: %T\n", response)

	// Always close response body (important)
	defer response.Body.Close()

	// Read complete response body as bytes
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Convert byte slice to string
	content := string(dataBytes)

	// Print small HTML content
	fmt.Println("\nResponse content:\n", content)
}
