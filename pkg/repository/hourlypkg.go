package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func CreateHourlyPackage(hourlyPackage models.HourlyPackage) error {
	if err := db.GetDBConn().Create(&hourlyPackage).Error; err != nil {
		logger.Error.Printf("[repository.CreateHourlyPackage] error creating hourly package: %v\n", err)
		return translateError(err)
	}
	return nil
}

func GetAllHourlyPackages() ([]models.HourlyPackage, error) {
	var hourlyPackages []models.HourlyPackage
	if err := db.GetDBConn().Find(&hourlyPackages).Error; err != nil {
		logger.Error.Printf("[repository.GetAllHourlyPackages] error fetching hourly packages: %v\n", err)
		return nil, translateError(err)
	}
	return hourlyPackages, nil
}

func GetHourlyPackageByID(id uint) (models.HourlyPackage, error) {
	var hourlyPackage models.HourlyPackage
	if err := db.GetDBConn().First(&hourlyPackage, id).Error; err != nil {
		logger.Error.Printf("[repository.GetHourlyPackageByID] error fetching hourly package by id: %v\n", err)
		return hourlyPackage, translateError(err)
	}
	return hourlyPackage, nil
}

func UpdateHourlyPackage(hourlyPackage models.HourlyPackage) error {
	if err := db.GetDBConn().Save(&hourlyPackage).Error; err != nil {
		logger.Error.Printf("[repository.UpdateHourlyPackage] error updating hourly package: %v\n", err)
		return translateError(err)
	}
	return nil
}

func DeleteHourlyPackage(id uint) error {
	if err := db.GetDBConn().Delete(&models.HourlyPackage{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeleteHourlyPackage] error deleting hourly package: %v\n", err)
		return translateError(err)
	}
	return nil
}
