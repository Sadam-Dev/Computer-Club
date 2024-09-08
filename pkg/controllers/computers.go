package controllers

import (
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAvailableComputers(c *gin.Context) {
	computers, err := service.GetAvailableComputers()

	if err != nil {
		handleError(c, err)
		return
	}

	if len(computers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Свободных мест в данный момент нет"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"computers": computers})
}
