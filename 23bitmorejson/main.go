package main

import (
	"encoding/json"
	"fmt"
)

/*
course struct
-------------
- Used to define JSON structure
- Struct tags control JSON keys and behavior
*/
type course struct {
	Name      string   `json:"coursename"` // Rename field in JSON
	Price     int      // Uses same name in JSON (price)
	Plateform string   `json:"plateform"`      // Custom JSON key
	Password  string   `json:"-"`              // Excluded from JSON output
	Tags      []string `json:"tags,omitempty"` // Omitted if nil or empty
}

func main() {
	fmt.Println("Welcome to JSON video")

	EncodeJson()
}

/*
EncodeJson
----------
- Creates course data
- Converts Go struct into JSON
- Prints formatted JSON output
*/
func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abcd123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "aaaa123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "bbbb123", nil},
	}

	// Package this data as JSON data
	// // This finalJson data without readable format
	// finalJson, err := json.Marshal(lcoCourses)

	// This finalJson data with readable format
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
