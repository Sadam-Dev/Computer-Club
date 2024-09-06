package models

import "time"

type BookingHistory struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	BookingID   uint      `json:"booking_id"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
