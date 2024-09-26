package controllers

import (
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRole(c *gin.Context) {
	roleIDStr := c.Param("id")                          // Получаем параметр как строку
	roleID, err := strconv.ParseUint(roleIDStr, 10, 32) // Преобразуем в uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := service.GetRoleByID(uint(roleID)) // Передаем
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.CreateRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create role"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role created successfully"})
}

func GetAllRoles(c *gin.Context) {
	roles, err := service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}
