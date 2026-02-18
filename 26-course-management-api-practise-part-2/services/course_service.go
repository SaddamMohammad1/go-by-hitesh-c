package services

import (
	"database/sql"
	"errors"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/repository"
)

func GetCourses(params models.CourseListParams) (models.PaginatedCoursesResponse, error) {
	// Defaults and caps
	if params.Limit <= 0 {
		params.Limit = 10
	}
	if params.Limit > 100 {
		params.Limit = 100
	}
	if params.Page <= 0 {
		params.Page = 1
	}

	data, err := repository.GetAllCourses(params)
	if err != nil {
		return models.PaginatedCoursesResponse{}, err
	}

	total, err := repository.CountCourses(params)
	if err != nil {
		return models.PaginatedCoursesResponse{}, err
	}

	var courses []models.CourseResponse
	for _, item := range data {
		courses = append(courses, models.CourseResponse{
			CourseId:    item.Course.CourseId,
			CourseName:  item.Course.CourseName,
			CoursePrice: item.Course.CoursePrice,
			Author:      item.Author,
		})
	}

	totalPage := (total + params.Limit - 1) / params.Limit
	if totalPage < 1 {
		totalPage = 1
	}

	return models.PaginatedCoursesResponse{
		Courses:   courses,
		Page:      params.Page,
		Limit:     params.Limit,
		Total:     total,
		TotalPage: totalPage,
	}, nil
}

func GetOneCourse(courseId int) (models.CourseResponse, error) {

	data, err := repository.GetOneCourse(courseId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CourseResponse{}, ErrCourseNotFound
		}
		return models.CourseResponse{}, err
	}

	return models.CourseResponse{
		CourseId:    data.Course.CourseId,
		CourseName:  data.Course.CourseName,
		CoursePrice: data.Course.CoursePrice,
		Author:      data.Author,
	}, nil
}

func CreateCourse(req models.CreateCourseRequest) (models.CourseResponse, error) {

	if req.CourseName == "" || req.CoursePrice <= 0 {
		return models.CourseResponse{}, ErrInvalidInput
	}

	course := models.Course{
		CourseName:  req.CourseName,
		CoursePrice: req.CoursePrice,
		AuthorID:    req.AuthorID,
	}

	createdCourse, err := repository.CreateCourse(course)
	if err != nil {
		return models.CourseResponse{}, err
	}

	// Optional: fetch author info for response
	author, err := repository.GetAuthorByID(req.AuthorID)
	if err != nil {
		return models.CourseResponse{}, err
	}

	return models.CourseResponse{
		CourseId:    createdCourse.CourseId,
		CourseName:  createdCourse.CourseName,
		CoursePrice: createdCourse.CoursePrice,
		Author:      author,
	}, nil
}

func UpdateCourse(id int, req models.UpdateCourseRequest) (models.CourseResponse, error) {

	if req.CourseName == "" || req.CoursePrice <= 0 {
		return models.CourseResponse{}, ErrInvalidInput
	}

	course := models.Course{
		CourseName:  req.CourseName,
		CoursePrice: req.CoursePrice,
		AuthorID:    req.AuthorID,
	}

	err := repository.UpdateCourse(id, course)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CourseResponse{}, ErrCourseNotFound
		}
		return models.CourseResponse{}, err
	}

	// Fetch updated data
	return GetOneCourse(id)
}

func DeleteCourse(id int) error {

	err := repository.DeleteCourse(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCourseNotFound
		}
		return err
	}

	return nil
}
