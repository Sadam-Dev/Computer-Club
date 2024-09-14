package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

func IsComputerAvailable(computerID uint, startTime, endTime time.Time) (bool, error) {
	var booking models.Booking
	err := db.GetDBConn().Where("computer_id = ? AND (start_time < ? AND end_time > ?)", computerID, endTime, startTime).First(&booking).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func CreateBooking(booking *models.Booking) error {
	err := db.GetDBConn().Create(booking).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCompletedBookings(currentTime time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := db.GetDBConn().Where("end_time <= ? AND is_completed = false", currentTime).Find(&bookings).Error
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func MarkBookingAsCompleted(bookingID uint) error {
	err := db.GetDBConn().Model(&models.Booking{}).Where("id = ?", bookingID).Update("is_completed", true).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateComputerAvailability(computerID uint, isAvailable bool) error {
	err := db.GetDBConn().Model(&models.Computer{}).Where("id = ?", computerID).Update("is_available", isAvailable).Error
	if err != nil {
		return err
	}
	return nil
}

func CompleteBooking(bookingID uint) error {
	var booking models.Booking
	err := db.GetDBConn().First(&booking, bookingID).Error
	if err != nil {
		logger.Error.Printf("[repository.CompleteBooking] Error finding booking: %v\n", err)
		return err
	}

	// Завершаем бронирование
	booking.IsCompleted = true
	err = db.GetDBConn().Save(&booking).Error
	if err != nil {
		logger.Error.Printf("[repository.CompleteBooking] Error completing booking: %v\n", err)
		return err
	}

	// Освобождаем компьютер
	var computer models.Computer
	err = db.GetDBConn().First(&computer, booking.ComputerID).Error
	if err != nil {
		logger.Error.Printf("[repository.CompleteBooking] Error finding computer: %v\n", err)
		return err
	}

	computer.IsAvailable = true
	err = db.GetDBConn().Save(&computer).Error
	if err != nil {
		logger.Error.Printf("[repository.CompleteBooking] Error freeing computer: %v\n", err)
		return err
	}

	return nil
}
