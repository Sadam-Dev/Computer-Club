package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"time"
)

func CreateComputer(computer models.Computer) error {
	return repository.CreateComputer(computer)
}

func GetAvailableComputers(startTime, endTime time.Time) ([]models.Computer, error) {
	computers, err := repository.GetAvailableComputers(startTime, endTime)
	if err != nil {
		return nil, errs.ErrValidationFailed
	}
	return computers, nil
}

func GetComputerByID(id uint) (models.Computer, error) {
	return repository.GetComputerByID(id)
}

func GetAllComputers() ([]models.Computer, error) {
	return repository.GetAllComputers()
}

func UpdateComputer(computer models.Computer) error {
	return repository.UpdateComputer(computer)
}

func DeleteComputer(id uint) error {
	return repository.DeleteComputer(id)
}
