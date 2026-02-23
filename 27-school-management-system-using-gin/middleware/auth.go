package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/SaddamMohammad1/27-school-management-system-using-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		// Check header exists
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header missing", nil)
			c.Abort()
			return
		}

		// Expect format: Bearer <token>
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization format", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", nil)
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token claims", nil)
			c.Abort()
			return
		}

		// Set values in context
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		// Continue
		c.Next()
	}
}
