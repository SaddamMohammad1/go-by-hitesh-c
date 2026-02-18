package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/utils"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware validates JWT and sets author_id and role in context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONResponse(w, http.StatusUnauthorized, false, nil, "missing authorization header")
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.JSONResponse(w, http.StatusUnauthorized, false, nil, "invalid authorization format")
			return
		}
		tokenStr := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default-secret-change-in-production"
		}
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			utils.JSONResponse(w, http.StatusUnauthorized, false, nil, "invalid or expired token")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.JSONResponse(w, http.StatusUnauthorized, false, nil, "invalid token claims")
			return
		}
		authorID, _ := claims["author_id"].(float64)
		role, _ := claims["role"].(string)
		ctx := context.WithValue(r.Context(), models.ContextKeyAuthorID, int(authorID))
		ctx = context.WithValue(ctx, models.ContextKeyRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// RequireAdmin wraps a handler and allows only admin role
func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(models.ContextKeyRole).(string)
		if role != models.RoleAdmin {
			utils.JSONResponse(w, http.StatusForbidden, false, nil, "admin access required")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireAuthorOrAdmin allows author or admin (for course create/update/delete and list own courses)
func RequireAuthorOrAdmin(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(models.ContextKeyRole).(string)
		if role != models.RoleAdmin && role != models.RoleAuthor {
			utils.JSONResponse(w, http.StatusForbidden, false, nil, "author or admin access required")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetAuthorIDFromContext returns author_id from context (0 if not set)
func GetAuthorIDFromContext(ctx context.Context) int {
	id, _ := ctx.Value(models.ContextKeyAuthorID).(int)
	return id
}

// GetRoleFromContext returns role from context
func GetRoleFromContext(ctx context.Context) string {
	role, _ := ctx.Value(models.ContextKeyRole).(string)
	return role
}

// OptionalAuthMiddleware parses JWT if present and sets context; does not return 401 if missing
func OptionalAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			next.ServeHTTP(w, r)
			return
		}
		tokenStr := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default-secret-change-in-production"
		}
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}
		authorID, _ := claims["author_id"].(float64)
		role, _ := claims["role"].(string)
		ctx := context.WithValue(r.Context(), models.ContextKeyAuthorID, int(authorID))
		ctx = context.WithValue(ctx, models.ContextKeyRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
