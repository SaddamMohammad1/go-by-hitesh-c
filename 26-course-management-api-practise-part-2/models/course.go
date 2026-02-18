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

// Pagination and filter for list courses
type CourseListParams struct {
	Page     int    // 1-based
	Limit    int    // default 10, max 100
	Search   string // search by course name (partial match)
	AuthorID int    // filter by author (0 = no filter)
	MinPrice int    // filter min price (0 = no filter)
	MaxPrice int    // filter max price (0 = no filter)
}

type PaginatedCoursesResponse struct {
	Courses   []CourseResponse `json:"courses"`
	Page      int              `json:"page"`
	Limit     int              `json:"limit"`
	Total     int              `json:"total"`
	TotalPage int              `json:"totalPage"`
}
