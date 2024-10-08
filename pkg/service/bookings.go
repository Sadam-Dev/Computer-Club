package service

import (
	"ComputerClub/logger"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"github.com/robfig/cron/v3"
)

//func CreateBooking(userID, computerID uint, startTime, endTime time.Time) (models.Booking, error) {
//	if endTime.Before(startTime) {
//		return models.Booking{}, fmt.Errorf("время окончания не может быть раньше времени начала")
//	}
//
//	available, err := repository.IsComputerAvailable(computerID, startTime, endTime)
//	if err != nil {
//		logger.Error.Printf("Error checking computer availability: %v\n", err)
//		return models.Booking{}, err
//	}
//
//	if !available {
//		return models.Booking{}, fmt.Errorf("компьютер уже забронирован на указанное время")
//	}
//
//	booking := models.Booking{
//		UserID:      userID,
//		ComputerID:  computerID,
//		StartTime:   startTime,
//		EndTime:     endTime,
//		IsCompleted: false,
//	}
//
//	err = repository.CreateBooking(&booking)
//	if err != nil {
//		logger.Error.Printf("Error creating booking: %v\n", err)
//		return models.Booking{}, err
//	}
//
//	err = repository.UpdateComputerAvailability(computerID, false)
//	if err != nil {
//		return models.Booking{}, fmt.Errorf("не удалось обновить статус компьютера: %v", err)
//	}
//
//	return booking, nil
//}
//

func CreateBooking(booking models.Booking) error {
	err := repository.CreateBooking(booking)
	if err != nil {
		return err
	}

	return nil
}

func GetBookingByID(id uint) (models.Booking, error) {
	return repository.GetBookingByID(id)
}

func GetAllBookings() ([]models.Booking, error) {
	return repository.GetAllBookings()
}

func UpdateBooking(booking models.Booking) error {
	return repository.UpdateBooking(booking)
}

func DeleteBooking(id uint) error {
	return repository.DeleteBooking(id)
}

func ProcessExpiredBookings() error {
	bookings, err := repository.GetExpiredBookings()
	if err != nil {
		return err
	}

	for _, booking := range bookings {
		// Завершение бронирования
		if err = repository.CompleteBooking(booking.ID); err != nil {
			return err
		}

		// Освобождение ПК
		if err = repository.SetComputerAvailable(booking.ComputerID); err != nil {
			return err
		}

		logger.Info.Printf("Booking %d completed and computer %d set as available\n", booking.ID, booking.ComputerID)
	}

	return nil
}

func StartBookingCleanupJob() {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		err := ProcessExpiredBookings()
		if err != nil {
			logger.Error.Printf("Error processing expired bookings: %v", err)
		}
	})
	c.Start()
}
