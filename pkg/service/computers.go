package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"time"
)

func CreateComputer(computer models.Computer) error {
	return repository.CreateComputer(computer)
}

func GetAvailableComputers(startTime, endTime time.Time) ([]models.Computer, error) {
	computers, err := repository.GetAvailableComputers(startTime, endTime)
	if err != nil {
		return nil, errs.ErrValidationFailed // Возможно, здесь вызывается ошибка
	}
	return computers, nil
}

func GetComputerByID(id uint) (models.Computer, error) {
	return repository.GetComputerByID(id)
}

func GetAllComputers() ([]models.Computer, error) {
	return repository.GetAllComputers()
}

func UpdateComputer(computer models.Computer) error {
	return repository.UpdateComputer(computer)
}

func DeleteComputer(id uint) error {
	return repository.DeleteComputer(id)
}

//func CompleteBooking(bookingID uint) error {
//	var booking models.Booking
//	err := db.GetDBConn().First(&booking, bookingID).Error
//	if err != nil {
//		return translateError(err)
//	}
//
//	// Проверяем, если текущее время больше или равно времени окончания бронирования
//	if time.Now().After(booking.EndTime) {
//		// Обновляем статус бронирования
//		booking.IsCompleted = true
//		err := db.GetDBConn().Save(&booking).Error
//		if err != nil {
//			return translateError(err)
//		}
//
//		// Освобождаем компьютер
//		err = repository.UpdateComputerAvailability(booking.ComputerID, true)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
