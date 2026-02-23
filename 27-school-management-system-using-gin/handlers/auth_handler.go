package handlers

import (
	"net/http"

	"github.com/SaddamMohammad1/27-school-management-system-using-gin/models"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/repository"
	"github.com/SaddamMohammad1/27-school-management-system-using-gin/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var user models.User

	// Bind JSON request
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, utils.MsgInvalidInput, err.Error())
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, utils.MsgServerError, "Failed to hash password")
		return
	}

	user.Password = string(hashedPassword)

	// Save user to DB
	if err := repository.CreateUser(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Registration failed", "Email already exists")
		return
	}

	// Return safe user data (no password)
	utils.SuccessResponse(c, http.StatusCreated, utils.MsgSuccessRegister, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

func Login(c *gin.Context) {

	var input models.User

	// Bind JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, utils.MsgInvalidInput, err.Error())
		return
	}

	// Get user from DB
	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, utils.MsgInvalidCreds, nil)
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, utils.MsgInvalidCreds, nil)
		return
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, utils.MsgServerError, "Token generation failed")
		return
	}

	// Only ONE response
	utils.SuccessResponse(c, http.StatusOK, utils.MsgSuccessLogin, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
