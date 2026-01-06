// This code see in 21webreqverbs section of Hitesh choudhry - class 30

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")

	// Calling function to perform POST request with formdata payload
	performPostFormRequest()
}

/*
performPostFormRequest
----------------------
- Creates form data
- Sends POST request using http.PostForm()
- Reads and prints server response
*/
func performPostFormRequest() {
	const myurl = "https://localhost:8000/postform"

	// fake formdata
	data := url.Values{}
	data.Add("firstname", "sad")
	data.Add("lastname", "hus")
	data.Add("email", "sad@gmail.com")

	// Here no need to pass data form type in parameter, because this PostForm method define to handle formdata request
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Read and print response from server
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
