package models

const (
	RoleAuthor = "author"
	RoleAdmin  = "admin"
)

type Author struct {
	ID           int    `json:"id"`
	Fullname     string `json:"fullname"`
	Website      string `json:"website"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"-"` // never expose in JSON
	Role         string `json:"role,omitempty"`
}

// CreateAuthorRequest used by admin to add a new author
type CreateAuthorRequest struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest for POST /login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthorResponse for API responses (no password)
type AuthorResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// LoginResponse returns JWT and user info
type LoginResponse struct {
	Token  string         `json:"token"`
	Author AuthorResponse `json:"author"`
}
