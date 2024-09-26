package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"errors"
	"fmt"
	"time"
)

func CreateHourlyPackage(hourlyPackage models.HourlyPackage) error {
	hourlyPackage.CreatedAt = time.Now()
	return repository.CreateHourlyPackage(hourlyPackage)
}

func GetAllHourlyPackages() ([]models.HourlyPackage, error) {
	return repository.GetAllHourlyPackages()
}

func GetHourlyPackageByID(id uint) (models.HourlyPackage, error) {
	return repository.GetHourlyPackageByID(id)
}

func UpdateHourlyPackage(hourlyPackage models.HourlyPackage) error {
	return repository.UpdateHourlyPackage(hourlyPackage)
}

func DeleteHourlyPackage(id uint) error {
	return repository.DeleteHourlyPackage(id)
}

func PurchaseHourlyPackage(username string, packageID uint) (models.Transaction, error) {
	// Получаем пакет
	hourlyPackage, err := repository.GetHourlyPackageByID(packageID)
	if err != nil {
		return models.Transaction{}, err
	}

	// Получаем баланс пользователя
	userBalance, err := repository.GetUserBalanceByUsername(username)
	if err != nil {
		return models.Transaction{}, err
	}

	// Проверяем, достаточно ли средств
	if userBalance.Balance < hourlyPackage.Price {
		return models.Transaction{}, errors.New("недостаточно средств на счете")
	}

	// Обновляем баланс
	newBalance := userBalance.Balance - hourlyPackage.Price

	fmt.Println(newBalance)

	// Обновляем баланс пользователя и игнорируем возвращаемый userBalance
	_, err = repository.UpdateUserBalanceByUsername(username, -hourlyPackage.Price)
	if err != nil {
		return models.Transaction{}, err
	}

	// Создаем транзакцию
	transaction := models.Transaction{
		UserID: userBalance.UserID,
		Amount: hourlyPackage.Price,
		Status: "completed",
	}

	if err = repository.CreateTransaction(transaction); err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}
