package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"errors"
)

func CreatePriceList(priceList models.PriceList) error {
	if priceList.PricePerHour <= 0 {
		return errs.ErrValidationFailed
	}

	err := repository.CreatePriceList(priceList)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPriceLists() ([]models.PriceList, error) {
	priceLists, err := repository.GetAllPriceLists()
	if err != nil {
		return nil, err
	}
	return priceLists, nil
}

func GetPriceListByID(id uint) (models.PriceList, error) {
	priceList, err := repository.GetPriceListByID(id)
	if err != nil {
		return models.PriceList{}, err
	}
	return priceList, nil
}

func UpdatePriceList(priceList models.PriceList) error {
	if priceList.PricePerHour <= 0 {
		return errs.ErrValidationFailed
	}

	return repository.UpdatePriceList(priceList)
}

func DeletePriceList(id uint) error {
	return repository.DeletePriceList(id)
}

func PurchaseTime(username string, categoryID uint, computerType string, hours int) (models.Transaction, error) {
	// Получаем цену за час
	price, err := repository.GetPriceByCategoryAndType(categoryID, computerType)
	if err != nil {
		return models.Transaction{}, err
	}

	// Рассчитываем общую стоимость
	totalCost := price.PricePerHour * float64(hours)

	// Получаем баланс пользователя
	userBalance, err := repository.GetUserBalanceByUsername(username)
	if err != nil {
		return models.Transaction{}, err
	}

	// Проверяем, достаточно ли средств
	if userBalance.Balance < totalCost {
		return models.Transaction{}, errors.New("недостаточно средств на счете")
	}

	// Обновляем баланс
	_, err = repository.UpdateUserBalanceByUsername(username, -totalCost)
	if err != nil {
		return models.Transaction{}, err
	}

	// Создаем транзакцию
	transaction := models.Transaction{
		UserID: userBalance.UserID,
		Amount: totalCost,
		Status: "completed",
	}

	if err = repository.CreateTransaction(transaction); err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}
