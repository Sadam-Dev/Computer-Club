package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.BindJSON(&booking); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrUnauthorized)
	}

	booking.UserID = userID

	// Логика создания бронирования
	if err := service.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Бронирование создано успешно", "booking": booking})
}

func GetBookingByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	booking, err := service.GetBookingByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, booking)
}

func GetAllBookings(c *gin.Context) {
	bookings, err := service.GetAllBookings()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func UpdateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.BindJSON(&booking); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	booking.ID = uint(id)

	if err := service.UpdateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking updated successfully"})
}

func DeleteBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteBooking(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "booking deleted successfully"})
}
