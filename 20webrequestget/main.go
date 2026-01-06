// This code see in 21webreqverbs section of Hitesh choudhry - class 28

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
	Topic: HTTP GET Request in Go

	There are TWO ways to read and print response body:
	1️⃣ Direct Read & Print (Simple)
	2️⃣ Using strings.Builder (Efficient & Recommended)

	Common Points:
	- http.Get() sends GET request
	- response.Body holds response stream
	- defer response.Body.Close() is mandatory
*/

func main() {

	/*
		Program execution starts from main()
	*/

	fmt.Println("Welcome to web verb video - LCO")

	performGetRequest() // Call GET API function
}

func performGetRequest() {

	/*
		Local API URL
		Backend can be written in Java / Python / PHP / Node / Django
	*/
	const myurl = "https://localhost:8000/get"

	/*
		Send HTTP GET request
	*/
	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		Always close response body
		to avoid memory leak
	*/
	defer response.Body.Close()

	/*
		Response information
	*/
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	/* -------------------------------------------------
		1️⃣ FIRST WAY: Direct Read and Print
		- Simple & easy
		- Suitable for small responses
	--------------------------------------------------*/
	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println("---- First Way Output ----")
	// fmt.Println(content)          // Prints byte slice
	// fmt.Println(string(content))  // Convert bytes to string

	/* -------------------------------------------------
		2️⃣ SECOND WAY: Using strings.Builder
		- Memory efficient
		- Recommended for large responses
	--------------------------------------------------*/

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("---- Second Way Output ----")
	fmt.Println("Byte Count is:", byteCount)
	fmt.Println(responseString.String())
}
