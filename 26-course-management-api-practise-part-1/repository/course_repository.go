package repository

import (
	"database/sql"

	"github.com/SaddamMohammad1/26-course-management-api-practise/config"
	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
)

type CourseWithAuthor struct {
	Course models.Course
	Author models.Author
}

// Get all courses with their authors from the database and return them as a list of CourseWithAuthor structs
func GetAllCourses() ([]CourseWithAuthor, error) {

	query := `
	SELECT c.course_id, c.course_name, c.course_price,
	       a.id, a.fullname, a.website
	FROM courses c
	LEFT JOIN authors a ON c.author_id = a.id
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []CourseWithAuthor

	for rows.Next() {

		var item CourseWithAuthor

		err := rows.Scan(
			&item.Course.CourseId,
			&item.Course.CourseName,
			&item.Course.CoursePrice,
			&item.Author.ID,
			&item.Author.Fullname,
			&item.Author.Website,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, item)

		// print the struct in a readable format
		// fmt.Printf("\nSTRUCT FORMAT: %+v\n", item)

		// Output
		// {
		// 	"Course": {
		// 		"CourseId": 1,
		// 		"CourseName": "Go Programming Masterclass",
		// 		"CoursePrice": 4999,
		// 		"AuthorID": 0
		// 	},
		// 	"Author": {
		// 		"ID": 1,
		// 		"Fullname": "Hitesh Choudhary",
		// 		"Website": "https://hitesh.ai"
		// 	}
		// }

	}

	return list, nil
}

// Get one course by ID with its author from the database and return it as a CourseWithAuthor struct
func GetOneCourse(courseId int) (CourseWithAuthor, error) {

	query := `
	SELECT c.course_id, c.course_name, c.course_price,
	       a.id, a.fullname, a.website
	FROM courses c
	LEFT JOIN authors a ON c.author_id = a.id
	WHERE c.course_id = $1
	`

	row := config.DB.QueryRow(query, courseId)

	var item CourseWithAuthor

	err := row.Scan(
		&item.Course.CourseId,
		&item.Course.CourseName,
		&item.Course.CoursePrice,
		&item.Author.ID,
		&item.Author.Fullname,
		&item.Author.Website,
	)

	if err != nil {
		return CourseWithAuthor{}, err
	}

	return item, nil
}

func CreateCourse(course models.Course) (models.Course, error) {

	query := `
		INSERT INTO courses (course_name, course_price, author_id)
		VALUES ($1, $2, $3)
		RETURNING course_id
	`

	err := config.DB.QueryRow(
		query,
		course.CourseName,
		course.CoursePrice,
		course.AuthorID,
	).Scan(&course.CourseId)

	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func UpdateCourse(id int, course models.Course) error {

	query := `
	UPDATE courses
	SET course_name = $1,
	    course_price = $2,
	    author_id = $3
	WHERE course_id = $4
	`

	result, err := config.DB.Exec(
		query,
		course.CourseName,
		course.CoursePrice,
		course.AuthorID,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteCourse(id int) error {
	query := `DELETE FROM courses WHERE course_id = $1`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
