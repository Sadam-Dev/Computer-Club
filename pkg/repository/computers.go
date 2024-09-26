package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
	"time"
)

func CreateComputer(computer models.Computer) error {
	if err := db.GetDBConn().Create(&computer).Error; err != nil {
		logger.Error.Printf("[repository.CreateComputer] error creating computer: %v\n", err)
		return translateError(err)
	}
	return nil
}

func GetAvailableComputers(startTime, endTime time.Time) ([]models.Computer, error) {
	var availableComputers []models.Computer
	if err := db.GetDBConn().Where("id NOT IN (?)",
		db.GetDBConn().Model(&models.Booking{}).
			Select("computer_id").
			Where("start_time < ? AND end_time > ?", endTime, startTime)).Find(&availableComputers).Error; err != nil {
		return nil, translateError(err)
	}
	return availableComputers, nil

}

func GetBookedComputers(startTime, endTime time.Time) ([]models.Computer, error) {
	var bookedComputers []models.Computer
	if err := db.GetDBConn().Model(&models.Booking{}).
		Select("computers.*").
		Joins("JOIN computers ON computers.id = bookings.computer_id").
		Where("bookings.start_time < ? AND bookings.end_time > ?", endTime, startTime).
		Find(&bookedComputers).Error; err != nil {
		return nil, err
	}
	return bookedComputers, nil
}

func GetComputerByID(id uint) (models.Computer, error) {
	var computer models.Computer
	if err := db.GetDBConn().First(&computer, id).Error; err != nil {
		return models.Computer{}, translateError(err)
	}
	return computer, nil
}

func GetAllComputers() ([]models.Computer, error) {
	var computers []models.Computer
	if err := db.GetDBConn().Find(&computers).Error; err != nil {
		return nil, translateError(err)
	}
	return computers, nil
}

func UpdateComputer(computer models.Computer) error {
	if err := db.GetDBConn().Save(&computer).Error; err != nil {
		logger.Error.Printf("[repository.UpdateComputer] error updating computer: %v\n", err)
		return translateError(err)
	}
	return nil
}

func DeleteComputer(id uint) error {
	if err := db.GetDBConn().Delete(&models.Computer{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeleteComputer] error deleting computer: %v\n", err)
		return translateError(err)
	}
	return nil
}
