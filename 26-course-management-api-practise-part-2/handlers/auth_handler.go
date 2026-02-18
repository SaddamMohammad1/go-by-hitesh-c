package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
	"github.com/SaddamMohammad1/26-course-management-api-practise/services"
	"github.com/SaddamMohammad1/26-course-management-api-practise/utils"
)

// LoginHandler handles POST /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
		return
	}
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
		return
	}
	resp, err := services.Login(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCreds) {
			utils.JSONResponse(w, http.StatusUnauthorized, false, nil, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusInternalServerError, false, nil, "internal server error")
		return
	}
	utils.JSONResponse(w, http.StatusOK, true, resp, "")
}
