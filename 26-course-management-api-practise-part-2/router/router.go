package router

import (
	"github.com/SaddamMohammad1/26-course-management-api-practise/handlers"
	"github.com/SaddamMohammad1/26-course-management-api-practise/middleware"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Public
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Admin only: add author
	r.HandleFunc("/authors", middleware.RequireAdmin(handlers.CreateAuthorHandler)).Methods("POST")

	// Courses: list/get support optional auth (author sees only own when logged in)
	r.HandleFunc("/courses", middleware.OptionalAuthMiddleware(handlers.GetCoursesHandler)).Methods("GET")
	r.HandleFunc("/courses/{id}", middleware.OptionalAuthMiddleware(handlers.GetOneCourseHandler)).Methods("GET")

	// Courses: create/update/delete require author or admin
	r.HandleFunc("/courses", middleware.RequireAuthorOrAdmin(handlers.CreateCourseHandler)).Methods("POST")
	r.HandleFunc("/courses/{id}", middleware.RequireAuthorOrAdmin(handlers.UpdateCourseHandler)).Methods("PUT")
	r.HandleFunc("/courses/{id}", middleware.RequireAuthorOrAdmin(handlers.DeleteCourseHandler)).Methods("DELETE")

	return r
}
