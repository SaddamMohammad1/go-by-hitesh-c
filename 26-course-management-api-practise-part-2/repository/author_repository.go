package repository

import (
	"database/sql"

	"github.com/SaddamMohammad1/26-course-management-api-practise/config"
	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
)

func GetAuthorByID(id int) (models.Author, error) {
	query := `SELECT id, fullname, website, email, password_hash, role FROM authors WHERE id = $1`
	var author models.Author
	err := config.DB.QueryRow(query, id).Scan(
		&author.ID,
		&author.Fullname,
		&author.Website,
		&author.Email,
		&author.PasswordHash,
		&author.Role,
	)
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

// GetAuthorByEmail returns author by email for login
func GetAuthorByEmail(email string) (models.Author, error) {
	query := `SELECT id, fullname, website, email, password_hash, role FROM authors WHERE email = $1`
	var author models.Author
	err := config.DB.QueryRow(query, email).Scan(
		&author.ID,
		&author.Fullname,
		&author.Website,
		&author.Email,
		&author.PasswordHash,
		&author.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Author{}, err
		}
		return models.Author{}, err
	}
	return author, nil
}

// CreateAuthor inserts a new author (used by admin)
func CreateAuthor(author models.Author) (models.Author, error) {
	query := `
		INSERT INTO authors (fullname, website, email, password_hash, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := config.DB.QueryRow(
		query,
		author.Fullname,
		author.Website,
		author.Email,
		author.PasswordHash,
		author.Role,
	).Scan(&author.ID)
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}
