package models

type UserRole struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	User   User `json:"-" gorm:"foreignKey:UserID"`
	RoleID uint `json:"role_id"`
	Role   Role `json:"-" gorm:"foreignKey:RoleID"`
}
