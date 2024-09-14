package db

import "ComputerClub/models"

func Migrate() error {
	err := GetDBConn().AutoMigrate(
		models.User{},
		models.Role{},
		models.Computer{},
		models.Booking{})
	if err != nil {
		return err
	}
	return nil
}
