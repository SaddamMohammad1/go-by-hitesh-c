package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/SaddamMohammad1/26-course-management-api-practise/config"
	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
)

type CourseWithAuthor struct {
	Course models.Course
	Author models.Author
}

// buildWhereClause builds WHERE conditions and args for list query
func buildCourseListWhere(params models.CourseListParams) (where string, args []interface{}) {
	var conditions []string
	var argNum int
	if params.Search != "" {
		argNum++
		conditions = append(conditions, fmt.Sprintf("c.course_name ILIKE $%d", argNum))
		args = append(args, "%"+params.Search+"%")
	}
	if params.AuthorID > 0 {
		argNum++
		conditions = append(conditions, fmt.Sprintf("c.author_id = $%d", argNum))
		args = append(args, params.AuthorID)
	}
	if params.MinPrice > 0 {
		argNum++
		conditions = append(conditions, fmt.Sprintf("c.course_price >= $%d", argNum))
		args = append(args, params.MinPrice)
	}
	if params.MaxPrice > 0 {
		argNum++
		conditions = append(conditions, fmt.Sprintf("c.course_price <= $%d", argNum))
		args = append(args, params.MaxPrice)
	}
	if len(conditions) > 0 {
		where = " WHERE " + strings.Join(conditions, " AND ")
	}
	return where, args
}

// CountCourses returns total count for given filters (no pagination)
func CountCourses(params models.CourseListParams) (int, error) {
	where, args := buildCourseListWhere(params)
	query := `SELECT COUNT(*) FROM courses c LEFT JOIN authors a ON c.author_id = a.id` + where
	var total int
	err := config.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

// GetAllCourses returns paginated, filtered, and searched courses with authors
func GetAllCourses(params models.CourseListParams) ([]CourseWithAuthor, error) {
	where, args := buildCourseListWhere(params)
	offset := (params.Page - 1) * params.Limit
	args = append(args, params.Limit, offset)
	n := len(args)
	query := `
	SELECT c.course_id, c.course_name, c.course_price,
	       a.id, a.fullname, a.website
	FROM courses c
	LEFT JOIN authors a ON c.author_id = a.id
	` + where + fmt.Sprintf(`
	ORDER BY c.course_id
	LIMIT $%d OFFSET $%d`, n-1, n)

	rows, err := config.DB.Query(query, args...)
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
