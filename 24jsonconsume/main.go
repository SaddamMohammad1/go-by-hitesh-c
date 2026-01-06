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
	fmt.Println("JSON Consume in go lang")
	DecodeJson()
}

func DecodeJson() {
	// Suppose this data comes from the server or api's response
	jsonDataFromWeb := []byte(`
		{
			"coursename": "MERN Bootcamp",
			"Price": 199,
			"plateform": "LearnCodeOnline.in",
			"tags": ["full-stack", "js"]
        }
	`)

	var lcoCourse course

	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// Some cases where you just want to add data to key value pair
	var myOnlineData map[string]interface{} // Here in JSON data key is clear, keys comes string, but value is not clear so here interface using in the place of value in this map
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
