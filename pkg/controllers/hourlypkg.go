package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateHourlyPackage(c *gin.Context) {
	var hourlyPackage models.HourlyPackage
	if err := c.BindJSON(&hourlyPackage); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateHourlyPackage(hourlyPackage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "hourly package created successfully"})
}

func GetAllHourlyPackages(c *gin.Context) {
	hourlyPackages, err := service.GetAllHourlyPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hourly_packages": hourlyPackages})
}

func GetHourlyPackageByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	hourlyPackage, err := service.GetHourlyPackageByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hourlyPackage)
}

func UpdateHourlyPackage(c *gin.Context) {
	var hourlyPackage models.HourlyPackage
	if err := c.BindJSON(&hourlyPackage); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	hourlyPackage.ID = uint(id)

	if err := service.UpdateHourlyPackage(hourlyPackage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "hourly package updated successfully"})
}

func DeleteHourlyPackage(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := service.DeleteHourlyPackage(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "hourly package deleted successfully"})
}

func PurchaseHourlyPackage(c *gin.Context) {
	var request struct {
		Username  string `json:"username"`
		PackageID uint   `json:"package_id"`
	}

	if err := c.BindJSON(&request); err != nil {
		handleError(c, err)
		return
	}

	transaction, err := service.PurchaseHourlyPackage(request.Username, request.PackageID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, transaction)
}
