package models

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"unique; not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
