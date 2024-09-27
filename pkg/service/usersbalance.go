package service

import (
	"ComputerClub/db"
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
)

func TopUpUserBalance(username string, amount float64) (models.UserBalance, error) {
	return repository.UpdateUserBalanceByUsername(username, amount)
}

func DeleteUserBalance(userID uint) error {
	return repository.DeleteUserBalance(userID)
}

func BuyTimeWithPackage(userID, hourlyPackageID uint) error {
	// Получаем информацию о пакете
	var hourlyPackage models.HourlyPackage
	if err := db.GetDBConn().First(&hourlyPackage, hourlyPackageID).Error; err != nil {
		return err
	}

	// Получаем баланс пользователя
	userBalance, err := repository.GetUserBalanceByUserID(userID)
	if err != nil {
		return err
	}

	// Проверяем достаточно ли средств
	if userBalance.Balance < hourlyPackage.FinalPrice {
		return errs.ErrBalanceInsufficient
	}

	// Списываем средства с баланса
	newBalance := userBalance.Balance - hourlyPackage.FinalPrice
	if err = repository.UpdateUserBalance(userID, newBalance); err != nil {
		return err
	}

	// Здесь можно добавить логику для добавления времени на компьютер
	// Например, создать или обновить запись бронирования для пользователя

	return nil
}

func BuyTimePerHour(userID uint, hours int, categoryID uint) error {
	// Получаем информацию о почасовой ставке
	var priceList models.PriceList
	if err := db.GetDBConn().Where("category_id = ?", categoryID).First(&priceList).Error; err != nil {
		return err
	}

	totalCost := priceList.PricePerHour * float64(hours)

	// Получаем баланс пользователя
	userBalance, err := repository.GetUserBalanceByUserID(userID)
	if err != nil {
		return err
	}

	// Проверяем достаточно ли средств
	if userBalance.Balance < totalCost {
		return errs.ErrBalanceInsufficient
	}

	// Списываем средства с баланса
	newBalance := userBalance.Balance - totalCost
	if err = repository.UpdateUserBalance(userID, newBalance); err != nil {
		return err
	}

	// Логика для добавления времени на компьютер

	return nil
}
