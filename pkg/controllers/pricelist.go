package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatePriceList(c *gin.Context) {
	var priceList models.PriceList
	if err := c.BindJSON(&priceList); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreatePriceList(priceList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "price list created successfully"})
}

func GetAllPriceLists(c *gin.Context) {
	priceLists, err := service.GetAllPriceLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get price lists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prices": priceLists})
}

func GetPriceListByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	priceList, err := service.GetPriceListByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, priceList)
}

func UpdatePriceList(c *gin.Context) {
	var priceList models.PriceList
	if err := c.BindJSON(&priceList); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.UpdatePriceList(priceList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "price list updated successfully"})
}

func DeletePriceList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.DeletePriceList(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "price list deleted successfully"})
}

func PurchaseTime(c *gin.Context) {
	var request struct {
		Username     string `json:"username"`
		CategoryID   uint   `json:"category_id"`
		ComputerType string `json:"computer_type"`
		Hours        int    `json:"hours"`
	}

	if err := c.BindJSON(&request); err != nil {
		handleError(c, err)
		return
	}

	transaction, err := service.PurchaseTime(request.Username, request.CategoryID, request.ComputerType, request.Hours)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, transaction)
}
