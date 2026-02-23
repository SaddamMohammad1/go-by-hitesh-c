package middleware

import (
	"net/http"

	"github.com/SaddamMohammad1/27-school-management-system-using-gin/utils"
	"github.com/gin-gonic/gin"
)

// RequireRole allows only specific roles
func RequireRole(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {

		roleInterface, exists := c.Get("role")

		if !exists {
			utils.ErrorResponse(c, http.StatusForbidden, "Role not found in token", nil)
			c.Abort()
			return
		}

		userRole := roleInterface.(string)

		for _, allowedRole := range roles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		utils.ErrorResponse(c, http.StatusForbidden, "Access denied", nil)
		c.Abort()
	}
}
