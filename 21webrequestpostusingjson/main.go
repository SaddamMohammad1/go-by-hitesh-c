// This code see in 21webreqverbs section of Hitesh choudhry - class 29

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")

	// Calling function to perform POST request with JSON payload
	performPostJsonRequest()
}

/*
performPostJsonRequest()
-----------------------
This function:
1. Creates a JSON payload
2. Sends it to the server using HTTP POST method
3. Reads and prints the response from server
*/
func performPostJsonRequest() {
	// API endpoint (URL) where POST request will be sent
	// NOTE: This is a local backend server URL
	const myurl = "https://localhost:8000/post"

	/*
		Create fake JSON payload
		------------------------
		strings.NewReader converts JSON string into io.Reader
		io.Reader is required by http.Post() method
	*/
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with go lang",
			"price": 0,
			"plateform": "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	/*
		Read response data
		------------------
		ioutil.ReadAll reads entire response body
		and returns it as byte slice
	*/
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
