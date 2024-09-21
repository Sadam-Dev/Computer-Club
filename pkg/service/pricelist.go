package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
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
