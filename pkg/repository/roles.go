package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func GetRoleByID(id uint) (role models.Role, err error) {
	err = db.GetDBConn().First(&role, id).Error
	if err != nil {
		logger.Error.Printf("[repository.GetRoleByID] Error getting role by ID: %s\n", err)
		return role, translateError(err)
	}
	return role, nil
}

func CreateRole(role models.Role) error {
	err := db.GetDBConn().Create(&role).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateRole] Error creating role: %s\n", err)
		return translateError(err)
	}
	return nil
}

func GetAllRoles() (roles []models.Role, err error) {
	err = db.GetDBConn().Find(&roles).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllRoles] Error getting all roles: %s\n", err)
		return nil, translateError(err)
	}
	return roles, nil
}

func InitializeRoles() error {
	roles := []models.Role{
		{Code: "user", Name: "User"},
		{Code: "admin", Name: "Admin"},
		{Code: "superAdmin", Name: "Super Admin"},
	}

	for _, role := range roles {
		if err := CreateRole(role); err != nil {
			return err
		}
	}
	return nil
}
