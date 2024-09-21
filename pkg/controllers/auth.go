package controllers

import (
	"ComputerClub/logger"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp
// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "account info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("Error binding JSON: %v", err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("Attempt to create user: %+v", user)

	err := service.CreateUser(user)
	if err != nil {
		logger.Error.Printf("Error creating user: %v", err)
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newDefaultResponse("account created successfully"))
}

// SignIn
// @Summary SignIn
// @Tags auth
// @Description sign in to account
// @ID sign-in-to-account
// @Accept json
// @Produce json
// @Param input body models.SignInInput true "sign-in info"
// @Success 200 {object} accessTokenResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	accessToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, accessTokenResponse{accessToken})
}
