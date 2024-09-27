package repository

import (
	"ComputerClub/db"
	"ComputerClub/errs"
	"ComputerClub/logger"
	"ComputerClub/models"
	"fmt"
)

func GetUserBalanceByUsername(username string) (models.UserBalance, error) {
	var user models.User
	err := db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserBalanceByUsername] Error getting user by username: %s\n", username)
		return models.UserBalance{}, translateError(err)
	}

	var userBalance models.UserBalance
	err = db.GetDBConn().Where("user_id = ?", user.ID).First(&userBalance).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserBalanceByUsername] Error getting balance for user ID: %d\n", user.ID)
		return userBalance, translateError(err)
	}

	return userBalance, nil
}

func UpdateUserBalanceByUsername(username string, amount float64) (models.UserBalance, error) {
	if amount <= 0 {
		return models.UserBalance{}, fmt.Errorf("amount must be greater than zero")
	}

	var user models.User
	err := db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserBalanceByUsername] Error getting user by username: %s\n", username)
		return models.UserBalance{}, translateError(err)
	}

	var userBalance models.UserBalance
	err = db.GetDBConn().Where("user_id = ?", user.ID).First(&userBalance).Error
	if err != nil && err == errs.ErrRecordNotFound {
		// Если запись не найдена, создаем новую
		userBalance = models.UserBalance{
			UserID:  user.ID,
			Balance: amount,
		}
		err = db.GetDBConn().Create(&userBalance).Error
		if err != nil {
			logger.Error.Printf("[repository.UpdateUserBalanceByUsername] Error creating balance for user ID: %d\n", user.ID)
			return userBalance, translateError(err)
		}
	} else if err == nil {
		// Проверка на достаточный баланс
		if userBalance.Balance+amount < 0 {
			return userBalance, fmt.Errorf("insufficient funds, balance cannot be negative")
		}
		// Если запись найдена, обновляем баланс
		userBalance.Balance += amount
		err = db.GetDBConn().Save(&userBalance).Error
		if err != nil {
			logger.Error.Printf("[repository.UpdateUserBalanceByUsername] Error updating balance for user ID: %d\n", user.ID)
			return userBalance, translateError(err)
		}
	} else {
		// Обрабатываем прочие ошибки
		logger.Error.Printf("[repository.UpdateUserBalanceByUsername] Error getting balance for user ID: %d\n", user.ID)
		return userBalance, translateError(err)
	}

	return userBalance, nil
}

func DeleteUserBalance(userID uint) error {
	if err := db.GetDBConn().Where("user_id = ?", userID).Delete(&models.UserBalance{}).Error; err != nil {
		return translateError(err)
	}

	return nil
}

func GetUserBalanceByUserID(userID uint) (models.UserBalance, error) {
	var balance models.UserBalance
	if err := db.GetDBConn().Where("user_id = ?", userID).First(&balance).Error; err != nil {
		return models.UserBalance{}, translateError(err)
	}
	return balance, nil
}

func UpdateUserBalance(userID uint, newBalance float64) error {
	return db.GetDBConn().Model(&models.UserBalance{}).Where("user_id = ?", userID).Update("balance", newBalance).Error
}
