package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/reposirory"
	"ComputerClub/utils"
	"errors"
)

func CreateUser(user models.User) error {
	userFromDB, err := reposirory.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if userFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	user.Role.Code = "user"

	user.Password = utils.GenerateHash(user.Password)

	err = reposirory.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() (users []models.User, err error) {
	users, err = reposirory.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
