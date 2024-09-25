package controllers

import (
	"ComputerClub/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Получение баланса пользователя через ID

func GetUserBalance(c *gin.Context) {
	username := c.Param("username")

	// Используем username для получения баланса
	balance, err := repository.GetUserBalanceByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance.Balance})
}
