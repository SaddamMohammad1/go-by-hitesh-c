package models

import "time"

type User struct {
	ID        int64      `json:"id"`
	SchoolID  *int64     `json:"school_id,omitempty"` // nullable
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`    // never return password in response
	Role      string     `json:"role"` // admin, teacher, student
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
