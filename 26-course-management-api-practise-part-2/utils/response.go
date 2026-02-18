package utils

import (
	"encoding/json"
	"net/http"

	"github.com/SaddamMohammad1/26-course-management-api-practise/models"
)

func JSONResponse(w http.ResponseWriter, status int, success bool, data any, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := models.APIResponse{
		Success: success,
		Data:    data,
		Error:   errMsg,
	}

	json.NewEncoder(w).Encode(response)
}
