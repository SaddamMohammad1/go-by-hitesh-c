package router

import (
	"github.com/SaddamMohammad1/26-course-management-api-practise/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/courses", handlers.GetCoursesHandler).Methods("GET")           // Get all
	r.HandleFunc("/courses/{id}", handlers.GetOneCourseHandler).Methods("GET")    // Get one
	r.HandleFunc("/courses", handlers.CreateCourseHandler).Methods("POST")        // Create
	r.HandleFunc("/courses/{id}", handlers.UpdateCourseHandler).Methods("PUT")    // Update
	r.HandleFunc("/courses/{id}", handlers.DeleteCourseHandler).Methods("DELETE") // Delete

	return r
}
