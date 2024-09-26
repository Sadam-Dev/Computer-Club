package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"errors"
)

func GetRoleByID(id uint) (models.Role, error) {
	return repository.GetRoleByID(id)
}

func CreateRole(role models.Role) error {
	if role.Code == "" || role.Name == "" {
		return errors.New("role code and name are required")
	}
	return repository.CreateRole(role)
}

func GetAllRoles() ([]models.Role, error) {
	return repository.GetAllRoles()
}
