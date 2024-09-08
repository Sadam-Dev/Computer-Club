package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func GetAvailableComputers() (computers []models.Computer, err error) {
	err = db.GetDBConn().Where("is_available = ?", true).Find(&computers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAvailableComputers] Error getting computers from database: %v\n", err.Error())
		return nil, translateError(err)
	}

	return computers, nil
}
