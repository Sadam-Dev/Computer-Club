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

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		configs.AppSettings.PostgresParams.SSLMode,
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to DB")

	dbConn = db
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
