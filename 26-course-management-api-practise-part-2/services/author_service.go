package services

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuthorNotFound = errors.New("author not found")
	ErrInvalidCreds   = errors.New("invalid email or password")
	ErrEmailExists    = errors.New("email already registered")
)

// Login validates credentials and returns JWT + author info
func Login(req models.LoginRequest) (models.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return models.LoginResponse{}, ErrInvalidCreds
	}
	author, err := repository.GetAuthorByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.LoginResponse{}, ErrInvalidCreds
		}
		return models.LoginResponse{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(author.PasswordHash), []byte(req.Password)); err != nil {
		return models.LoginResponse{}, ErrInvalidCreds
	}
	token, err := generateJWT(author.ID, author.Email, author.Role)
	if err != nil {
		return models.LoginResponse{}, err
	}
	return models.LoginResponse{
		Token: token,
		Author: models.AuthorResponse{
			ID:       author.ID,
			Fullname: author.Fullname,
			Website:  author.Website,
			Email:    author.Email,
			Role:     author.Role,
		},
	}, nil
}

func generateJWT(authorID int, email, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-change-in-production"
	}
	claims := jwt.MapClaims{
		"author_id": authorID,
		"email":     email,
		"role":      role,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// CreateAuthor creates a new author (admin only). Password is hashed.
func CreateAuthor(req models.CreateAuthorRequest) (models.AuthorResponse, error) {
	if req.Fullname == "" || req.Email == "" || req.Password == "" {
		return models.AuthorResponse{}, ErrInvalidInput
	}
	_, err := repository.GetAuthorByEmail(req.Email)
	if err == nil {
		return models.AuthorResponse{}, ErrEmailExists
	}
	if err != nil && err != sql.ErrNoRows {
		return models.AuthorResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.AuthorResponse{}, err
	}
	author := models.Author{
		Fullname:     req.Fullname,
		Website:      req.Website,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         models.RoleAuthor,
	}
	created, err := repository.CreateAuthor(author)
	if err != nil {
		return models.AuthorResponse{}, err
	}
	return models.AuthorResponse{
		ID:       created.ID,
		Fullname: created.Fullname,
		Website:  created.Website,
		Email:    created.Email,
		Role:     created.Role,
	}, nil
}
