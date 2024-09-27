package models

import "time"

type Computer struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	IsAvailable bool       `json:"is_available" gorm:"default:true"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"-" gorm:"index"`
}

type SwagComputer struct {
	Name        string `json:"name"`         // Название компьютера
	Description string `json:"description"`  // Описание компьютера
	IsAvailable bool   `json:"is_available"` // Доступен ли компьютер
}
