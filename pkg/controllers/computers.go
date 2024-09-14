package controllers

import (
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBind(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateComputer(computer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create computer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"computer": "Computer created successfully"})

}

func GetAvailableComputers(c *gin.Context) {
	computers, err := service.GetAvailableComputers()

	if err != nil {
		handleError(c, err)
		return
	}

	fmt.Printf("Computers found: %v\n", computers)

	if len(computers) == 0 {
		c.JSON(http.StatusOK, gin.H{"computers": "Свободных мест в данный момент нет"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"computers": computers})
}
