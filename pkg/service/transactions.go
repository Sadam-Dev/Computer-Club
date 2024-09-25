package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
)

// AddFunds пополняет баланс пользователя и создает транзакцию
func AddFunds(username string, amount float64) (models.Transaction, error) {
	// Обновляем баланс пользователя через username
	updatedBalance, err := repository.UpdateUserBalanceByUsername(username, amount)
	if err != nil {
		return models.Transaction{}, err
	}

	// Создаем транзакцию
	transaction := models.Transaction{
		UserID: updatedBalance.UserID,
		Amount: amount,
		Status: "completed",
	}

	// Сохраняем транзакцию в базе данных
	if err = repository.CreateTransaction(transaction); err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}
