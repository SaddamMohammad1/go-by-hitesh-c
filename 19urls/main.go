package main

import (
	"fmt"
	"net/url"
)

// Sample URL with:
// scheme + host + port + path + query params
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghfs343"

func main() {

	// URL handling in Go
	fmt.Println("Welcome to handling urls in go lang")
	fmt.Println(myurl)

	// Parse string URL into structured URL object
	result, _ := url.Parse(myurl)

	// URL components
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=ghfs343

	// Extract query parameters as map[string][]string
	qparams := result.Query()

	// Type of query params
	fmt.Printf("The type of query params are: %T\n", qparams)

	// Access individual query values
	fmt.Println(qparams["coursename"]) // [reactjs]
	fmt.Println(qparams["paymentid"])  // [ghfs343]

	// Loop through all query params
	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	// Create URL manually using url.URL struct
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
		// NOTE: RawQuery is used for query params, not RawPath
		RawQuery: "user=sad",
	}

	// Convert URL struct to string
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
