package main

import (
	"log"
	"net/http"

	"github.com/SaddamMohammad1/26-course-management-api-practise/config"
	"github.com/SaddamMohammad1/26-course-management-api-practise/router"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Setup the router
	r := router.SetupRouter()

	// Start the server
	log.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
