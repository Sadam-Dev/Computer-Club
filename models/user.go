package models

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username" gorm:"unique; not null"`
	Password    string    `json:"password" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"unique; not null"`
	Email       string    `json:"email" gorm:"unique; not null"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
	IsBlocked   bool      `json:"is_blocked" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	RoleID      uint      `json:"role_id"`
	Role        Role      `json:"-" gorm:"foreignKey:RoleID"`
}

type SwagUser struct {
	FullName string `json:"full_name"`
	Username string `json:"username" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`
}
