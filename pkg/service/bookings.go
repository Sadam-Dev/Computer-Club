package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
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
