package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"errors"
)

func isCategoryNameUnique(name string) (bool, error) {
	category, err := repository.GetCategoryByName(name)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return false, err
	}
	return category.ID == 0, nil // Если ID 0, то категории не существует
}

func CreateCategory(category models.Category) error {
	isUnique, err := isCategoryNameUnique(category.Name)
	if err != nil {
		return err
	}
	if !isUnique {
		return errs.ErrCategoryNameExists
	}

	err = repository.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryByID(id uint) (models.Category, error) {
	return repository.GetCategoryByID(id)
}

func GetAllCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

func UpdateCategory(category models.Category) error {
	isUnique, err := isCategoryNameUnique(category.Name)
	if err != nil {
		return err
	}
	if !isUnique {
		return errs.ErrCategoryNameExists
	}
	return repository.UpdateCategory(category)
}

func DeleteCategory(id uint) error {
	return repository.DeleteCategory(id)
}
