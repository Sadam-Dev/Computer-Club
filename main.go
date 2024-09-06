package main

import (
	"ComputerClub/configs"
	"ComputerClub/db"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
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

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}
}
