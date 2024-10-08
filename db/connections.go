package db

import (
	"ComputerClub/configs"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbConn *gorm.DB

func ConnectToDB() error {

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s timezone=%s",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		configs.AppSettings.PostgresParams.SSLMode,
		os.Getenv("DB_PASSWORD"),
		configs.AppSettings.PostgresParams.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to DB")

	dbConn = db
	return nil
}

func CloseDBConn() error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
