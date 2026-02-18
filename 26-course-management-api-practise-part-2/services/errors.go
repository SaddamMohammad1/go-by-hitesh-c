package services

import "errors"

var (
	ErrCourseNotFound = errors.New("course not found")
	ErrInvalidID      = errors.New("invalid course id")
	ErrInvalidInput   = errors.New("invalid input data")
)
