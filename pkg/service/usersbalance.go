package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
)

func TopUpUserBalance(username string, amount float64) (models.UserBalance, error) {
	return repository.UpdateUserBalanceByUsername(username, amount)
}
