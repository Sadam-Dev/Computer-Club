package service

import (
	"ComputerClub/errs"
	"ComputerClub/pkg/reposirory"
	"ComputerClub/utils"
	"errors"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)

	user, err := reposirory.GetUserByUsernameAndPassword(username, password)

	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return "", errs.ErrIncorrectUsernameOrPassword
		}
		return "", err
	}

	accessToken, err = GenerateToken(user.ID, user.Username, user.Role.Code)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
