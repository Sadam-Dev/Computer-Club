package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/logger"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser
// @Summary Register a new user
// @Tags users
// @Description Register a new user (only Admin can do this)
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "User Information"
// @Success 201 {string} DefaultResponse "User created successfully!!!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users [post]
// @Security ApiKeyAuth
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("[CreateUser] Invalid input: %v, IP: [%s]\n", err, c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("[CreateUser] IP: [%s] attempting to create user: %s\n", c.ClientIP(), user.Username)

	if err := service.CreateUser(user); err != nil {
		logger.Error.Printf("[CreateUser] Error creating user: %v, IP: [%s]\n", err, c.ClientIP())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!!!"})
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users or a single user by username
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        username  query   string  false  "Username to filter the user"
// @Success      200  {array}   models.SwagUser  "List of users"
// @Success      200  {object}  models.User  "Single user details"
// @Failure      403  {object}  ErrorResponse "Access Denied"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /user [get]
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// GetUserByID godoc
// @Summary Retrieve a user by ID
// @Description Get user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] Invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] Error retrieving user: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserByID godoc
// @Summary Update a user by ID
// @Description Update user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUserByID] Invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var updateUser models.User
	if err = c.BindJSON(&updateUser); err != nil {
		logger.Error.Printf("[controllers.UpdateUserByID] Error binding JSON: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	user, err := service.UpdateUserByID(uint(id), updateUser)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUserByID] Error updating user: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"user":    user,
	})
}

// DeleteUserByID godoc
// @Summary Delete a user by ID
// @Description Delete a user by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUserByID] Invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := service.DeleteUserByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUserByID] Failed to delete user: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	logger.Info.Printf("User with ID [%d] successfully deleted by client with IP [%s]\n", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"user":    user,
	})
}
