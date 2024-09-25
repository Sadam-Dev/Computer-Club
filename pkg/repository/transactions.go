package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

// Создание транзакции
func CreateTransaction(transaction models.Transaction) error {
	err := db.GetDBConn().Create(&transaction).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateTransaction] Error creating transaction: %v\n", err)
		return translateError(err)
	}

	return nil
}
