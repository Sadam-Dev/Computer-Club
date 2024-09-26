package service

import (
	"ComputerClub/errs"
	"ComputerClub/logger"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"ComputerClub/utils"
	"errors"
)

func CreateUser(user models.User) error {
	// Проверяем уникальность username
	if existingUser, err := repository.GetUserByUsername(user.Username); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		logger.Error.Printf("[service.CreateUser] error checking username uniqueness: %v\n", err)
		return err
	} else if existingUser.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	// Назначаем стандартную роль пользователю (например, "user")
	user.RoleID = 1

	// Хэшируем пароль перед сохранением
	user.Password = utils.GenerateHash(user.Password)

	// Создаем пользователя и баланс
	err := repository.CreateUser(&user)
	if err != nil {
		logger.Error.Printf("[service.CreateUser] error creating user: %v\n", err)
		return err
	}

	return nil
}

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		logger.Error.Printf("[service.GetAllUsers] error getting all users: %v\n", err)
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		logger.Error.Printf("[service.GetUserByID] error getting user by ID: %v\n", err)
		return user, err
	}

	return user, nil
}

func UpdateUserByID(id uint, updateUser models.User) (user models.User, err error) {
	user, err = repository.UpdateUserByID(id, updateUser)
	if err != nil {
		logger.Error.Printf("[service.UpdateUserByID] error updating user by ID: %v\n", err)
		return user, err
	}

	return user, nil
}

func DeleteUserByID(id uint) (user models.User, err error) {
	user, err = repository.DeleteUserByID(id)
	if err != nil {
		logger.Error.Printf("[service.DeleteUserByID] error deleting user by ID: %v\n", err)
		return user, err
	}

	return user, nil
}
