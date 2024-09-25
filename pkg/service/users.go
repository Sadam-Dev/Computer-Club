package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"ComputerClub/utils"
	"errors"
)

func CreateUser(user models.User) error {
	// Проверяем уникальность username
	if existingUser, err := repository.GetUserByUsername(user.Username); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	} else if existingUser.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	// Назначаем стандартную роль пользователю (например, "user")
	user.RoleID = 1

	// Хэшируем пароль перед сохранением
	user.Password = utils.GenerateHash(user.Password)

	// Создаем пользователя и баланс
	err := repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUserByID(id uint, updateUser models.User) (user models.User, err error) {
	user, err = repository.UpdateUserByID(id, updateUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUserByID(id uint) (user models.User, err error) {
	user, err = repository.DeleteUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
