package repository

import (
	"github.com/SaddamMohammad1/26-course-management-api-practise/config"
	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
)

func GetAuthorByID(id int) (models.Author, error) {

	query := `SELECT id, fullname, website FROM authors WHERE id = $1`

	var author models.Author

	err := config.DB.QueryRow(query, id).Scan(
		&author.ID,
		&author.Fullname,
		&author.Website,
	)

	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}
