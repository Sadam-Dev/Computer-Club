package controllers

import (
	"ComputerClub/db"
	"ComputerClub/models"
	"ComputerClub/pkg/repository"
	"ComputerClub/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateComputer godoc
// @Summary      Создать новый компьютер
// @Description  Создает запись о новом компьютере в базе данных
// @Tags         computers
// @Accept       json
// @Produce      json
// @Param        computer  body      models.Computer  true  "Данные компьютера"
// @Success      201       {object}  gin.H "Computer created successfully"
// @Failure      400       {object}  map[string]string "error"
// @Router       /computers [post]
func CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBind(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateComputer(computer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create computer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"computer": "Computer created successfully"})
}

// GetAvailableComputers godoc
// @Summary      Получить список доступных компьютеров
// @Description  Возвращает список всех доступных (свободных) компьютеров
// @Tags         computers
// @Produce      json
// @Success      200  {array}   models.Computer  "Список доступных компьютеров"
// @Failure      400  {object}  map[string]string "error"
// @Router       /computers/available [get]
func GetAvailableComputers(c *gin.Context) {
	computers, err := service.GetAvailableComputers(true)

	if err != nil {
		handleError(c, err)
		return
	}

	fmt.Printf("Computers found: %v\n", computers)

	if len(computers) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "Свободных мест в данный момент нет"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"computers": computers})
}

// GetBookingComputers godoc
// @Summary      Получить список не доступных компьютеров
// @Description  Возвращает список всех не доступных (не свободных) компьютеров
// @Tags         computers
// @Produce      json
// @Success      200  {array}   models.Computer  "Список не оступных компьютеров"
// @Failure      400  {object}  map[string]string "error"
// @Router       /computers/available [get]
func GetBookingComputers(c *gin.Context) {
	computers, err := service.GetAvailableComputers(false)
	repository.UpdateComputerAvailabilityAuto(db.GetDBConn())
	if err != nil {
		handleError(c, err)
		return
	}

	fmt.Printf("Computers found: %v\n", computers)

	if len(computers) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "Свободных мест в данный момент нет"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"computers": computers})
}
