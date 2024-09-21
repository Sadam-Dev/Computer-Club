package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func CreatePriceList(priceList models.PriceList) error {
	if err := db.GetDBConn().Create(&priceList).Error; err != nil {
		logger.Error.Printf("[repository.CreatePriceList] error creating price list: %v\n", err)
		return translateError(err)
	}
	return nil
}

func GetPriceListByID(id uint) (models.PriceList, error) {
	var priceList models.PriceList
	if err := db.GetDBConn().First(&priceList, id).Error; err != nil {
		logger.Error.Printf("[repository.GetPriceListByID] error fetching price list by ID: %v\n", err)
		return priceList, translateError(err)
	}
	return priceList, nil
}

func GetAllPriceLists() ([]models.PriceList, error) {
	var priceLists []models.PriceList
	if err := db.GetDBConn().Find(&priceLists).Error; err != nil {
		logger.Error.Printf("[repository.GetAllPriceLists] error retrieving price lists: %v\n", err)
		return nil, translateError(err)
	}
	return priceLists, nil
}

func UpdatePriceList(priceList models.PriceList) error {
	if err := db.GetDBConn().Save(&priceList).Error; err != nil {
		logger.Error.Printf("[repository.UpdatePriceList] error updating price list: %v\n", err)
		return translateError(err)
	}
	return nil
}

func DeletePriceList(id uint) error {
	if err := db.GetDBConn().Delete(&models.PriceList{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeletePriceList] error deleting price list: %v\n", err)
		return translateError(err)
	}
	return nil
}
