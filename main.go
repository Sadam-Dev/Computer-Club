package main

import (
	"ComputerClub/configs"
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/pkg/controllers"
	"ComputerClub/pkg/service"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Error loading .env file. Error is %s", err)))
	}

	err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}

	err = logger.Init()
	if err != nil {
		panic(err)
	}

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	err = controllers.RunRoutes()
	if err != nil {
		panic(err)
	}

	service.StartUpdatingComputerAvailability(1 * time.Minute)

}
