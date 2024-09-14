package service

import (
	"ComputerClub/logger"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"fmt"
	"time"
)

func CreateBooking(userID, computerID uint, startTime, endTime time.Time) (models.Booking, error) {
	if endTime.Before(startTime) {
		return models.Booking{}, fmt.Errorf("время окончания не может быть раньше времени начала")
	}

	available, err := repository.IsComputerAvailable(computerID, startTime, endTime)
	if err != nil {
		logger.Error.Printf("Error checking computer availability: %v\n", err)
		return models.Booking{}, err
	}

	if !available {
		return models.Booking{}, fmt.Errorf("компьютер уже забронирован на указанное время")
	}

	booking := models.Booking{
		UserID:      userID,
		ComputerID:  computerID,
		StartTime:   startTime,
		EndTime:     endTime,
		IsCompleted: false,
	}

	err = repository.CreateBooking(&booking)
	if err != nil {
		logger.Error.Printf("Error creating booking: %v\n", err)
		return models.Booking{}, err
	}

	err = repository.UpdateComputerAvailability(computerID, false)
	if err != nil {
		return models.Booking{}, fmt.Errorf("не удалось обновить статус компьютера: %v", err)
	}

	return booking, nil
}

func UpdateComputerAvailabilityAfterBooking() error {
	currentTime := time.Now()

	completedBookings, err := repository.GetCompletedBookings(currentTime)
	if err != nil {
		logger.Error.Printf("Ошибка получения завершенных бронирований: %v\n", err)
		return err
	}

	for _, booking := range completedBookings {
		err = repository.UpdateComputerAvailability(booking.ComputerID, true)
		if err != nil {
			logger.Error.Printf("Ошибка обновления статуса компьютера: %v\n", err)
			return err
		}

		err = repository.MarkBookingAsCompleted(booking.ID)
		if err != nil {
			logger.Error.Printf("Ошибка обновления статуса бронирования: %v\n", err)
			return err
		}
	}

	return nil
}

func StartUpdatingComputerAvailability(interval time.Duration) {
	go func() {
		for {
			err := UpdateComputerAvailabilityAfterBooking()
			if err != nil {
				fmt.Println("Ошибка при обновлении доступности компьютеров:", err)
			}

			time.Sleep(interval)
		}
	}()
}

func StartBookingTimer(bookingID uint, endTime time.Time) {
	duration := time.Until(endTime)

	// Запускаем таймер в новой горутине
	go func() {
		time.Sleep(duration) // Ждем окончания времени бронирования
		err := repository.CompleteBooking(bookingID)
		if err != nil {
			// Логируем ошибку, если возникла проблема при завершении бронирования
			logger.Error.Printf("[service.StartBookingTimer] Error completing booking: %v\n", err)
		} else {
			// Возможно, можно выполнить действие для выключения компьютера
			logger.Info.Printf("Booking %d completed and computer freed.\n", bookingID)
		}
	}()
}
