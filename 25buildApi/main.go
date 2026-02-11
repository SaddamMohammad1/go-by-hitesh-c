package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file name
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Initialize slice of course - DB seed
var courses []Course

// Middleware, helper functions - file name
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "" // Here only check these two if these two empty then return true otherwise false
	return c.CourseName == "" // Here only check course name if course name empty then return true otherwise false
}

func main() {
	fmt.Println("API - Learning Go")
	r := mux.NewRouter()

	// Seeding data
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "ReactJS Bootcamp",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Saddam",
			Website:  "https://react.dev",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "NextJS Bootcamp",
		CoursePrice: 599,
		Author: &Author{
			Fullname: "Saddam",
			Website:  "https://next.dev",
		},
	})

	// Routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	// Listen to port 8000 and start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Controllers - file name

// Serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to our API</h1>")) // Here retun in bytes because w.Write() accepts byte slice as argument and we are converting string to byte slice using []byte() conversion
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

	// Behind the scenes, the json.NewEncoder(w).Encode(courses) is doing the following:
	// 1. It creates a new JSON encoder that writes to the http.ResponseWriter (w).
	// 2. It takes the courses slice and encodes it into JSON format.
	// 3. The resulting JSON is then written to the response body, which is sent back to the client.

	// encoder := json.NewEncoder(w)
	// jsonBytes := json.Marshal(courses)
	// w.Write(jsonBytes)
	// w.Write([]byte("\n"))

}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	// Grab id from request
	params := mux.Vars(r) // Here we are using mux.Vars() function to extract the variables from the URL path. It returns a map of variable names to their corresponding values. In this case, we are extracting the "id" variable from the URL and storing it in the params map.

	// Example: If the request URL is /course/123, then params will be a map with a single key "id" and its value will be "123". So, params["id"] will give us the value "123", which we can use to find the course with that specific ID in our courses slice.

	// Loop through courses and find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// Return response when id not found
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// What if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// What about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// Check only if title is dublicate
	// loop and title match if match then return dublicate error message
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exist")
			return
		}
	}

	// Generate unique id, string
	// Append course into courses
	course.CourseId = strconv.Itoa(rand.IntN(100)) // Generate a random number between 0 and 100 and convert it to string and assign it to course id
	courses = append(courses, course)

	// Return the course which is just created
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	// Grad id from request
	params := mux.Vars(r)
	id := params["id"]

	// Decode updated data ONCE
	var updatedCourse Course
	if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	// Check duplicate course name (ignore same ID)
	for _, c := range courses {
		if c.CourseName == updatedCourse.CourseName && c.CourseId != id {
			json.NewEncoder(w).Encode("Course name already exists")
			return
		}
	}

	// Loop through courses, find matching id and remove the course from courses and add the updated course to courses
	for index, existingCourse := range courses {
		if existingCourse.CourseId == id {
			// Remove old course
			courses = append(courses[:index], courses[index+1:]...)
			// Keep same ID
			updatedCourse.CourseId = id
			// Add updated course
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	// Return response when id not found
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// Grad id from request
	params := mux.Vars(r)
	// loop id through courses, find matching id and remove the course from courses
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove the course from courses
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	// Return response when id not found
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
