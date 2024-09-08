package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
)

func GetAvailableComputers() (computers []models.Computer, err error) {
	computers, err = repository.GetAvailableComputers()
	if err != nil {
		return nil, err
	}

	return computers, nil
}
