package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/SaddamMohammad1/26-course-management-api-practise/middleware"
	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/services"
	"github.com/SaddamMohammad1/26-course-management-api-practise/utils"
	"github.com/gorilla/mux"
)

func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	params := parseCourseListParams(r)
	// Author sees only their courses; admin/public see by params
	if authorID := middleware.GetAuthorIDFromContext(r.Context()); authorID != 0 {
		if middleware.GetRoleFromContext(r.Context()) == models.RoleAuthor {
			params.AuthorID = authorID
		}
	}
	courses, err := services.GetCourses(params)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	utils.JSONResponse(w, http.StatusOK, true, courses, "")
}

func parseCourseListParams(r *http.Request) models.CourseListParams {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))
	authorID, _ := strconv.Atoi(q.Get("author_id"))
	minPrice, _ := strconv.Atoi(q.Get("min_price"))
	maxPrice, _ := strconv.Atoi(q.Get("max_price"))
	if limit <= 0 {
		limit = 10
	}
	return models.CourseListParams{
		Page:     page,
		Limit:    limit,
		Search:   q.Get("search"),
		AuthorID: authorID,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}
}

func GetOneCourseHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, services.ErrInvalidID.Error())
		return
	}
	course, err := services.GetOneCourse(id)
	if err != nil {
		if errors.Is(err, services.ErrCourseNotFound) {
			utils.JSONResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	// Author can only view their own course
	if role := middleware.GetRoleFromContext(r.Context()); role == models.RoleAuthor {
		if middleware.GetAuthorIDFromContext(r.Context()) != course.Author.ID {
			utils.JSONResponse(w, http.StatusForbidden, false, nil, "forbidden: not your course")
			return
		}
	}
	utils.JSONResponse(w, http.StatusOK, true, course, "")
}

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
		return
	}
	// Author can only create under their own id; admin can set any author_id
	if middleware.GetRoleFromContext(r.Context()) == models.RoleAuthor {
		req.AuthorID = middleware.GetAuthorIDFromContext(r.Context())
	}
	course, err := services.CreateCourse(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidInput) {
			utils.JSONResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	utils.JSONResponse(w, http.StatusCreated, true, course, "")
}

func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, services.ErrInvalidID.Error())
		return
	}
	// Author can only update their own course
	if middleware.GetRoleFromContext(r.Context()) == models.RoleAuthor {
		course, _ := services.GetOneCourse(id)
		if course.Author.ID != middleware.GetAuthorIDFromContext(r.Context()) {
			utils.JSONResponse(w, http.StatusForbidden, false, nil, "forbidden: not your course")
			return
		}
	}
	var req models.UpdateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
		return
	}
	course, err := services.UpdateCourse(id, req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidInput) {
			utils.JSONResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		if errors.Is(err, services.ErrCourseNotFound) {
			utils.JSONResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	utils.JSONResponse(w, http.StatusOK, true, course, "")
}

func DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, services.ErrInvalidID.Error())
		return
	}
	// Author can only delete their own course
	if middleware.GetRoleFromContext(r.Context()) == models.RoleAuthor {
		course, _ := services.GetOneCourse(id)
		if course.Author.ID != middleware.GetAuthorIDFromContext(r.Context()) {
			utils.JSONResponse(w, http.StatusForbidden, false, nil, "forbidden: not your course")
			return
		}
	}
	err = services.DeleteCourse(id)
	if err != nil {
		if errors.Is(err, services.ErrCourseNotFound) {
			utils.JSONResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	utils.JSONResponse(w, http.StatusOK, true, "course deleted successfully", "")
}
