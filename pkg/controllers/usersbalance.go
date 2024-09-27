package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/pkg/repository"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserBalance(c *gin.Context) {
	username := c.Param("username")

	balance, err := repository.GetUserBalanceByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance.Balance})
}

func DeleteUserBalance(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrUnauthorized)
		return
	}

	err := service.DeleteUserBalance(userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User balance deleted successfully"})
}

func BuyTimeWithPackage(c *gin.Context) {
	userID := c.GetUint(userIDCtx) // Получаем ID пользователя из контекста
	hourlyPackageID, err := strconv.ParseUint(c.Param("package_id"), 10, 64)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.BuyTimeWithPackage(userID, uint(hourlyPackageID)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Time purchased successfully using package"})
}

func BuyTimePerHour(c *gin.Context) {
	userID := c.GetUint(userIDCtx) // Получаем ID пользователя из контекста
	categoryID, err := strconv.ParseUint(c.Param("category_id"), 10, 64)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	hours, err := strconv.Atoi(c.Query("hours"))
	if err != nil || hours <= 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.BuyTimePerHour(userID, hours, uint(categoryID)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Time purchased successfully by hours"})
}
