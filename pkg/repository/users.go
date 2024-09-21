package repository

import (
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/models"
)

func getUserByID(id uint) (models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.getUserByID] Error getting user by ID: %d\n", id)
		return user, translateError(err)
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] erroe creating user: %v\n", err)
		return translateError(err)

	}
	return nil
}

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] Error getting all users: %s\n", err.Error())
		return nil, translateError(err)
	}

	return users, nil
}

func GetUserByID(id uint) (models.User, error) {
	return getUserByID(id)
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] Error getting user by Username: %s\n", err)
		return user, translateError(err)
	}

	return user, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] Error getting user by Username and Password: %s\n", err)
		return user, translateError(err)
	}

	return user, nil
}

func UpdateUserByID(id uint, updateUser models.User) (models.User, error) {
	user, err := getUserByID(id)
	if err != nil {
		return user, err
	}

	err = db.GetDBConn().Model(&user).Updates(updateUser).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserByID] Error updating user by ID: %s\n", err)
		return user, translateError(err)
	}

	return user, nil

}

func DeleteUserByID(id uint) (models.User, error) {
	user, err := getUserByID(id)

	err = db.GetDBConn().Delete(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUserByID] Error deleting user by ID: %d\n", id)
		return user, translateError(err)
	}

	return user, nil
}
