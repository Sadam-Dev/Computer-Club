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

func GetAvailableComputers() (computers []models.Computer, err error) {
	computers, err = repository.GetAvailableComputers()
	if err != nil {

		return nil, err
	}

	return computers, nil
}
