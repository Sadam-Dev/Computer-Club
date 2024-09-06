package models

type UserRole struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id" gorm:"foreignKey:UserID"`
	RoleID uint `json:"role_id" gorm:"foreignKey:RoleID"`
}
