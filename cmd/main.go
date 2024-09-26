package main

import (
	"ComputerClub/configs"
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/pkg/controllers"
	"ComputerClub/pkg/service"
	"ComputerClub/server"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Computer Club API
// @version 1.0
// @description API Server for Computer Club Application

// @host localhost:8181
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		logger.Error.Fatalf("Error loading the .env file: %s", err)
	}
	fmt.Println("Environment variables loaded successfully!")

	// Чтение настроек из файла конфигурации
	if err := configs.ReadSettings(); err != nil {
		logger.Error.Fatalf("Error reading settings: %s", err)
	}
	fmt.Println("Configuration loaded successfully!")

	// Загружаем временную зону "Asia/Dushanbe"
	_, err := time.LoadLocation(configs.AppSettings.PostgresParams.TimeZone)
	if err != nil {
		fmt.Println("Ошибка при загрузке временной зоны:", err)
		return
	}

	// Инициализация логгера
	if err := logger.Init(); err != nil {
		logger.Error.Fatalf("Error initializing logger: %s", err)
	}
	fmt.Println("Logger initialized successfully!")

	// Подключение к базе данных
	if err := db.ConnectToDB(); err != nil {
		logger.Error.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.CloseDBConn(); err != nil {
			logger.Error.Printf("Error closing database connection: %v", err)
		}
	}()
	fmt.Println("Database connected successfully!")

	// Миграция базы данных
	if err := db.Migrate(); err != nil {
		logger.Error.Fatalf("Failed to run database migrations: %v", err)
	}
	fmt.Println("Database migrated successfully!")

	// Инициализация сервера
	mainServer := new(server.Server)

	service.StartBookingCleanupJob()

	// Запуск сервера в горутине
	go func() {
		if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			logger.Error.Fatalf("Error starting HTTP server: %s", err)
		}
	}()

	// Ожидание сигнала для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nStart of program completion\n")

	// Закрытие соединения с БД, если необходимо
	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			logger.Error.Fatalf("Error closing DB: %s", err)
		}
	} else {
		logger.Error.Fatalf("Error getting *sql.DB from GORM: %s", err)
	}
	fmt.Println("The connection to the DB was successfully closed")

	// Используем контекст с тайм-аутом для завершения работы сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := mainServer.Shutdown(ctx); err != nil {
		fmt.Println(err.Error())
		logger.Error.Fatalf("Error during server shutdown: %s", err)
	}
	fmt.Println("Server shut down gracefully!")

	fmt.Println("HTTP service successfully disabled")
	fmt.Println("End of program completion")
}
