package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func CreateComputer(computer models.Computer) error {
	err := db.GetDBConn().Create(&computer).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateComputer] Error adding computer to database: %v\n", err.Error())
		return err
	}
	return nil
}

func GetAvailableComputers(isAvailable bool) (computers []models.Computer, err error) {
	err = db.GetDBConn().Where("is_available = ?", isAvailable).Find(&computers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAvailableComputers] Error getting computers from database: %v\n", err.Error())
		return nil, translateError(err)
	}

	return computers, nil
}

// UpdateComputerAvailabilityAuto автоматически обновляет доступность компьютеров
func UpdateComputerAvailabilityAuto(db *gorm.DB) {
	//now := time.Now()
	//formattedTime := now.Format("2006-01-02 15:04")
	var bookings []models.Booking
	//fmt.Println(now)

	// Находим все бронирования, которые завершились и компьютеры ещё заняты
	if err := db.Where("end_time <=  NOW() + INTERVAL '5 hours' AND is_completed = ?", false).Find(&bookings).Error; err != nil {
		log.Println("Ошибка при получении завершённых бронирований:", err)
		return
	}
	fmt.Println("len(bookings):", len(bookings))

	for _, booking := range bookings {
		fmt.Println(booking.ComputerID)
		// Обновляем статус компьютера на доступный
		if err := db.Model(&models.Computer{}).Where("id = ? AND is_available = false", booking.ComputerID).Update("is_available", true).Error; err != nil {
			fmt.Println("Ошибка при обновлении статуса компьютера:", err)
		} else {
			fmt.Printf("Компьютер с ID %d теперь доступен\n", booking.ComputerID)
			booking.IsCompleted = true
			if err := db.Save(&booking).Error; err != nil {
				fmt.Println("error saving computer to database:", err.Error())
			}
		}
	}
	fmt.Println("UpdateComputerAvailabilityAuto")
}
