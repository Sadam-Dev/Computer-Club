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

// CreateComputer godoc
// @Summary      Create a new computer
// @Description  Creates a new computer entry in the system.
// @Tags         computers
// @Accept       json
// @Produce      json
// @Param        computer  body      models.Computer  true  "Computer Details"
// @Success      201       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /computers [post]
func CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.BindJSON(&computer); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateComputer(computer); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "computer created successfully"})
}

// GetAvailableComputers godoc
// @Summary      Get available computers
// @Description  Retrieves all available computers that are not booked.
// @Tags         computers
// @Produce      json
// @Success      200  {array}   models.Computer
// @Failure      500  {object}  map[string]string
// @Router       /computers/available [get]
func GetAvailableComputers(c *gin.Context) {
	availableComputers, err := service.GetAvailableComputers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"computers": availableComputers})
}

// GetComputerByID godoc
// @Summary      Get computer by ID
// @Description  Retrieves details of a computer by its ID.
// @Tags         computers
// @Produce      json
// @Param        id   path      int  true  "Computer ID"
// @Success      200  {object}  models.Computer
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /computers/{id} [get]
func GetComputerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	computer, err := service.GetComputerByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, computer)
}

// GetBookedComputers godoc
// @Summary      Get booked computers
// @Description  Retrieves computers that are currently booked.
// @Tags         computers
// @Produce      json
// @Success      200  {array}   models.Computer
// @Failure      500  {object}  map[string]string
// @Router       /computers/booked [get]
func GetBookedComputers(c *gin.Context) {
	currentTime := time.Now()

	bookedComputers, err := service.GetBookedComputers(currentTime)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, bookedComputers)
}

// GetAllComputers godoc
// @Summary      Get all computers
// @Description  Retrieves all computers in the system.
// @Tags         computers
// @Produce      json
// @Success      200  {array}   models.Computer
// @Failure      500  {object}  map[string]string
// @Router       /computers [get]
func GetAllComputers(c *gin.Context) {
	computers, err := service.GetAllComputers()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, computers)
}

// UpdateComputer godoc
// @Summary      Update a computer
// @Description  Updates the details of an existing computer by its ID.
// @Tags         computers
// @Accept       json
// @Produce      json
// @Param        id        path      int             true  "Computer ID"
// @Param        computer  body      models.Computer  true  "Updated Computer Details"
// @Success      200       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Failure      404       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /computers/{id} [put]
func UpdateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.BindJSON(&computer); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	computer.ID = uint(id)

	if err := service.UpdateComputer(computer); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "computer updated successfully"})
}

// DeleteComputer godoc
// @Summary      Delete a computer
// @Description  Deletes a computer by its ID.
// @Tags         computers
// @Produce      json
// @Param        id   path      int  true  "Computer ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /computers/{id} [delete]
func DeleteComputer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteComputer(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "computer deleted successfully"})
}
