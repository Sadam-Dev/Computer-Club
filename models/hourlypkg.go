package models

import "time"

type HourlyPackage struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CategoryID    uint      `json:"category_id" gorm:"not null"`
	ComputerType  string    `json:"computer_type" gorm:"not null"`
	Hours         int       `json:"hours" gorm:"not null"`
	Price         float64   `json:"price" gorm:"not null"`
	Discount      float64   `json:"discount" gorm:"default:0"`
	FinalPrice    float64   `json:"final_price"`
	EffectiveDate time.Time `json:"effective_date" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
