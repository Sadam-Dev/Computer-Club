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

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.UpdateUserByID(uint(id), updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"user":    user,
	})
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Error.Printf("Client with IP [%s] successfully deleted\n", c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"user":    user,
	})
}
