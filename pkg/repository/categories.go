package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func CreateCategory(category models.Category) (err error) {
	if err = db.GetDBConn().Create(&category).Error; err != nil {
		logger.Error.Printf("[repository.CreateCategory] error creating category: %v\n", err)
		return translateError(err)
	}
	return nil
}

func GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := db.GetDBConn().First(&category, id).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCategoryByID] error getting category: %v\n", err)
		return category, translateError(err)
	}
	return category, nil
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := db.GetDBConn().Find(&categories).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllCategories] error getting categories: %v\n", err)
		return nil, translateError(err)
	}
	return categories, nil
}

func UpdateCategory(category models.Category) error {
	if err := db.GetDBConn().Save(&category).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCategory] error updating category: %v\n", err)
		return translateError(err)
	}
	return nil
}

func DeleteCategory(id uint) error {
	if err := db.GetDBConn().Delete(&models.Category{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCategory] error deleting category: %v\n", err)
		return translateError(err)
	}
	return nil
}

func GetCategoryByName(name string) (models.Category, error) {
	var category models.Category
	err := db.GetDBConn().Where("name = ?", name).First(&category).Error
	if err != nil {
		return category, translateError(err)
	}
	return category, nil
}
