package db

import "ComputerClub/models"

func Migrate() error {
	err := GetDBConn().AutoMigrate(
		models.User{},
		models.Role{},
		models.Computer{},
		models.Booking{},
		models.Role{},
		models.PriceList{},
		models.HourlyPackage{},
		models.Category{},
		models.UserBalance{},
		models.Transaction{})
	if err != nil {
		return err
	}
	return nil
}
