package repository

import (
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/config"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/models"
)

// Create new user
func CreateUser(user *models.User) error {

	query := `
	INSERT INTO users (school_id, name, email, password, role)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at;
	`

	err := config.DB.QueryRow(
		query,
		user.SchoolID,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt)

	return err
}

// Get user by email (for login)
func GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	query := `
	SELECT id, school_id, name, email, password, role, is_active, created_at
	FROM users
	WHERE email = $1;
	`

	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.SchoolID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
