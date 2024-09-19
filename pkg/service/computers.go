package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
)

func CreateComputer(computer models.Computer) error {
	err := repository.CreateComputer(computer)
	if err != nil {
		return err
	}
	return nil
}

func GetAvailableComputers(isAvailable bool) (computers []models.Computer, err error) {
	computers, err = repository.GetAvailableComputers(isAvailable)
	if err != nil {

		return nil, err
	}

	return computers, nil
}
