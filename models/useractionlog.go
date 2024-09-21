package models

import "time"

type UserActionLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Action    string    `json:"action" gorm:"not null"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
