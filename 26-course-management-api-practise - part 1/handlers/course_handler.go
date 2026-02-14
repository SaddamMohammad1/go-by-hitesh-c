package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/services"
	"github.com/SaddamMohammad1/26-course-management-api-practise/utils"
	"github.com/gorilla/mux"
)

func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {

	courses, err := services.GetCourses()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, courses, "")
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

	utils.JSONResponse(w, http.StatusOK, true, course, "")
}

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {

	var req models.CreateCourseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
		return
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
