package main

import (
	"ComputerClub/configs"
	"ComputerClub/db"
	"ComputerClub/logger"
	"ComputerClub/pkg/controllers"
	"ComputerClub/pkg/repository"
	"ComputerClub/server"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title ComputerClub API
// @version 1.02
// @description API Server for ComputerClub Application

// @host localhost:8181
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загрузка .env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	// Чтение настроек
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	// Инициализация логгера
	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}

	// Подключение к базе данных
	var err error
	err = db.ConnectToDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	// Миграция базы данных
	if err = db.Migrate(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}

	// Настройка cron-задания
	c := cron.New()
	_, err = c.AddFunc("@every 30s", func() {
		repository.UpdateComputerAvailabilityAuto(db.GetDBConn())
	})
	if err != nil {
		log.Fatal("Ошибка при создании cron-задания:", err)
	}

	// Запуск cron
	go func() {
		c.Start()
	}()
	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	// Ожидание сигнала для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nНачало завершения программ\n")

	// Закрытие соединения с БД, если необходимо
	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println("Соединение с БД успешно закрыто")

	// Используем контекст с тайм-аутом для завершения работы сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = mainServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")
}
