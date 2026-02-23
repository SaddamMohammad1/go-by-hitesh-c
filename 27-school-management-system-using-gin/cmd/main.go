package main

import (
	"net/http"

	"github.com/SaddamMohammad1/27-school-management-system-using-gin/config"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/handlers"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/middleware"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/utils"
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

	// Protected Routes Group
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {

		userID, _ := c.Get("user_id")
		role, _ := c.Get("role")

		utils.SuccessResponse(c, http.StatusOK, "Protected route accessed", gin.H{
			"user_id": userID,
			"role":    role,
		})
	})

	// Admin only route
	protected.GET("/admin-dashboard",
		middleware.RequireRole("admin"),
		func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Welcome Admin", nil)
		},
	)

	// Teacher only route
	protected.GET("/teacher-dashboard",
		middleware.RequireRole("teacher"),
		func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Welcome Teacher", nil)
		},
	)

	// Student only route
	protected.GET("/student-dashboard",
		middleware.RequireRole("student"),
		func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Welcome Student", nil)
		},
	)

	r.Run(":8080")
}
