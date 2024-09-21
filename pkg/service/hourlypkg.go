package service

import (
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"time"
)

func CreateHourlyPackage(hourlyPackage models.HourlyPackage) error {
	hourlyPackage.CreatedAt = time.Now()
	return repository.CreateHourlyPackage(hourlyPackage)
}

func GetAllHourlyPackages() ([]models.HourlyPackage, error) {
	return repository.GetAllHourlyPackages()
}

func GetHourlyPackageByID(id uint) (models.HourlyPackage, error) {
	return repository.GetHourlyPackageByID(id)
}

func UpdateHourlyPackage(hourlyPackage models.HourlyPackage) error {
	return repository.UpdateHourlyPackage(hourlyPackage)
}

func DeleteHourlyPackage(id uint) error {
	return repository.DeleteHourlyPackage(id)
}
