package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

// getUserByID retrieves a user by ID from the database.
func getUserByID(id uint) (models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.getUserByID] error getting user by ID: %d\n", id)
		return user, translateError(err)
	}
	logger.Info.Printf("[repository.getUserByID] user retrieved successfully with ID: %d\n", id)
	return user, nil
}

// CreateUser creates a new user and initializes their balance.
func CreateUser(user *models.User) error {
	if err := db.GetDBConn().Create(user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return translateError(err)
	}

	// Создание записи баланса для пользователя
	userBalance := models.UserBalance{
		UserID:  user.ID,
		Balance: 0,
	}
	if err := db.GetDBConn().Create(&userBalance).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user balance: %v\n", err)
		return translateError(err)
	}

	logger.Info.Printf("[repository.CreateUser] user created successfully with ID: %d\n", user.ID)
	return nil
}

// GetAllUsers retrieves all users from the database, excluding deleted ones.
func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %v\n", err)
		return nil, translateError(err)
	}

	logger.Info.Printf("[repository.GetAllUsers] users retrieved successfully\n")
	return users, nil
}

// GetUserByID retrieves a user by ID.
func GetUserByID(id uint) (models.User, error) {
	return getUserByID(id)
}

// GetUserByUsername retrieves a user by username.
func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %s\n", err)
		return user, translateError(err)
	}

	logger.Info.Printf("[repository.GetUserByUsername] user retrieved successfully with username: %s\n", username)
	return user, nil
}

// GetUserByUsernameAndPassword retrieves a user by username and password.
func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().
		Preload("Role"). // Загрузка связанной роли
		Where("username = ? AND password = ?", username, password).
		First(&user).Error

	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %s\n", err)
		return user, translateError(err)
	}

	return user, nil
}

// UpdateUserByID updates a user by ID.
func UpdateUserByID(id uint, updateUser models.User) (models.User, error) {
	user, err := getUserByID(id)
	if err != nil {
		return user, err
	}

	err = db.GetDBConn().Model(&user).Updates(updateUser).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserByID] error updating user by ID: %s\n", err)
		return user, translateError(err)
	}

	logger.Info.Printf("[repository.UpdateUserByID] user updated successfully with ID: %d\n", id)
	return user, nil
}

// DeleteUserByID deletes a user by ID (soft delete).
func DeleteUserByID(id uint) (models.User, error) {
	user, err := getUserByID(id)
	if err != nil {
		return user, err
	}

	err = db.GetDBConn().Delete(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUserByID] error deleting user by ID: %d\n", id)
		return user, translateError(err)
	}

	logger.Info.Printf("[repository.DeleteUserByID] user deleted successfully with ID: %d\n", id)
	return user, nil
}

// HardDeleteUserByID permanently deletes a user by ID.
func HardDeleteUserByID(id uint) (models.User, error) {
	user, err := getUserByID(id)
	if err != nil {
		return user, err
	}

	err = db.GetDBConn().Unscoped().Delete(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.HardDeleteUserByID] error deleting user by ID: %v\n", err)
		return user, translateError(err)
	}
	logger.Info.Printf("[repository.HardDeleteUserByID] user hard deleted successfully with ID: %d\n", id)
	return user, nil
}
