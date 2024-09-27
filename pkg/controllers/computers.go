package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.BindJSON(&computer); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateComputer(computer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "computer created successfully"})
}

func GetAvailableComputers(c *gin.Context) {

	availableComputers, err := service.GetAvailableComputers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"computers": availableComputers})
}

func GetComputerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	computer, err := service.GetComputerByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, computer)
}

func GetBookedComputersHandler(c *gin.Context) {
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	if len(startTimeStr) == 10 {
		startTimeStr += "T00:00:00Z"
	}
	if len(endTimeStr) == 10 {
		endTimeStr += "T23:59:59Z"
	}

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}
	// Вызов функции из слоя service
	bookedComputers, err := service.GetBookedComputers(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookedComputers)
}

func GetAllComputers(c *gin.Context) {
	computers, err := service.GetAllComputers()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, computers)
}

func UpdateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.BindJSON(&computer); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	computer.ID = uint(id)

	if err := service.UpdateComputer(computer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "computer updated successfully"})
}

func DeleteComputer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteComputer(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "computer deleted successfully"})
}
