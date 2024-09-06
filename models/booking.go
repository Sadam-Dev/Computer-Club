package models

import "time"

type Booking struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UserID      uint      `json:"user_id"`
	ComputerID  uint      `json:"computer_id"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
