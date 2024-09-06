package models

import "time"

type Transaction struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	BookingID uint      `json:"booking_id"`
	Amount    float64   `json:"amount" gorm:"not null"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}
