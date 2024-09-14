package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func CreateComputer(computer models.Computer) error {
	err := db.GetDBConn().Create(&computer).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateComputer] Error adding computer to database: %v\n", err.Error())
		return err
	}
	return nil
}

func GetAvailableComputers() (computers []models.Computer, err error) {
	err = db.GetDBConn().Where("is_available = ?", true).Find(&computers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAvailableComputers] Error getting computers from database: %v\n", err.Error())
		return nil, translateError(err)
	}

	return computers, nil
}
