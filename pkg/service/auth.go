package service

import (
	"ComputerClub/errs"
	"ComputerClub/pkg/repository"
	"ComputerClub/utils"
	"errors"
	"fmt"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)

	user, err := repository.GetUserByUsernameAndPassword(username, password)

	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return "", errs.ErrIncorrectUsernameOrPassword
		}
		return "", err
	}

	role := user.Role.Code

	// Логируем роль для отладки
	fmt.Printf("Extracted role for user %s: %s\n", user.Username, role)

	accessToken, err = GenerateToken(user.ID, user.Username, user.Role.Code)
	if err != nil {
		return "", err
	}

	fmt.Printf("Generated token for user %s with role %s: %s\n", user.Username, user.Role.Code, accessToken)

	return accessToken, nil
}
