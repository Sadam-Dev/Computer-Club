package models

import "time"

type PriceList struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CategoryID    uint      `json:"category_id" gorm:"not null"`
	ComputerType  string    `json:"computer_type" gorm:"not null"`
	PricePerHour  float64   `json:"price_per_hour" gorm:"not null"`
	EffectiveDate time.Time `json:"effective_date" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
