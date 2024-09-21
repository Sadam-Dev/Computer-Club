package models

import "time"

type ProfitReport struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TotalProfit float64   `json:"total_profit" gorm:"not null"`
	ReportDate  time.Time `json:"report_date" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
