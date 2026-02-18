package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/services"
	"github.com/SaddamMohammad1/26-course-management-api-practise/utils"
)

// CreateAuthorHandler creates a new author (admin only)
func CreateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
		return
	}
	author, err := services.CreateAuthor(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidInput) {
			utils.JSONResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		if errors.Is(err, services.ErrEmailExists) {
			utils.JSONResponse(w, http.StatusConflict, false, nil, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	utils.JSONResponse(w, http.StatusCreated, true, author, "")
}
