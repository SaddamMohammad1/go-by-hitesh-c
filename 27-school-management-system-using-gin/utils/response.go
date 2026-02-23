package utils

import (
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/models"
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, status int, message string, data any) {
	response := models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(status, response)
}

func ErrorResponse(c *gin.Context, status int, message string, err any) {
	response := models.APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(status, response)
}
