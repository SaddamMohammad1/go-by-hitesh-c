package main

import (
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/config"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to database
	config.ConnectDB()

	// Create Gin router
	r := gin.Default()

	// Test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "School Management System Running",
		})
	})

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
