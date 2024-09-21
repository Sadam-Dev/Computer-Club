package models

import "time"

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`          // Ссылка на пользователя
	Amount    float64   `json:"amount" gorm:"not null"`           // Сумма транзакции
	Status    string    `json:"status"`                           // Статус (например, "completed", "pending")
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Время создания транзакции
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Время обновления транзакции
}
