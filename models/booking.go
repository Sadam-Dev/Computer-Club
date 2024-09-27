package models

import "time"

type Booking struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	ComputerID  uint      `json:"computer_id" gorm:"not null"`
	Computer    Computer  `json:"-" gorm:"foreignKey:ComputerID"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SwagBooking struct {
	UserID      uint      `json:"user_id"`      // ID пользователя
	ComputerID  uint      `json:"computer_id"`  // ID компьютера
	StartTime   time.Time `json:"start_time"`   // Время начала брони
	EndTime     time.Time `json:"end_time"`     // Время окончания брони
	IsCompleted bool      `json:"is_completed"` // Завершена ли бронь
}
