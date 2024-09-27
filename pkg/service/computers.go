package service

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
)

func CreateComputer(computer models.Computer) error {
	return repository.CreateComputer(computer)
}

func GetAvailableComputers() ([]models.Computer, error) {
	computers, err := repository.GetAvailableComputers()
	if err != nil {
		return nil, errs.ErrValidationFailed
	}
	return computers, nil
}

// GetBookedComputers - логика получения всех забронированных компьютеров
func GetBookedComputers() ([]models.Computer, error) {
	// Вызов функции из repository для получения забронированных ПК
	bookedComputers, err := repository.GetBookedComputers()
	if err != nil {
		return nil, err
	}

	return bookedComputers, nil
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
