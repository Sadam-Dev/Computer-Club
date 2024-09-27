package models

import "time"

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique; not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index"`
}

type SwagCategory struct {
	Name        string `json:"name"`        // Название категории
	Description string `json:"description"` // Описание категории
}
