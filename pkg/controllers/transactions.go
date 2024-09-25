package controllers

import (
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Пополнение баланса через username

func TopUpBalanceByUsername(c *gin.Context) {
	type TopUpRequest struct {
		Username string  `json:"username"`
		Amount   float64 `json:"amount"`
	}

	var req TopUpRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Пополнение баланса через username
	transaction, err := service.AddFunds(req.Username, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем созданную транзакцию
	c.JSON(http.StatusOK, transaction)
}
