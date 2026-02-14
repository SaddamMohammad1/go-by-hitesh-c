package models

type Course struct {
	CourseId    int
	CourseName  string
	CoursePrice int
	AuthorID    int
}

type CourseResponse struct {
	CourseId    int    `json:"courseId"`
	CourseName  string `json:"courseName"`
	CoursePrice int    `json:"coursePrice"`
	Author      Author `json:"author"` // Here Author struct access because course.go and author.go are in same package(models)
}

type CreateCourseRequest struct {
	CourseName  string `json:"courseName"`
	CoursePrice int    `json:"coursePrice"`
	AuthorID    int    `json:"authorId"`
}

type UpdateCourseRequest struct {
	CourseName  string `json:"courseName"`
	CoursePrice int    `json:"coursePrice"`
	AuthorID    int    `json:"authorId"`
}
