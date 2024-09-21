package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func CreateBooking(c *gin.Context) {
//	var booking models.Booking
//
//	if err := c.ShouldBindJSON(&booking); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	createdBooking, err := service.CreateBooking(booking.UserID, booking.ComputerID, booking.StartTime, booking.EndTime)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Бронирование успешно создано", "booking": createdBooking})
//}

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.BindJSON(&booking); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "booking created successfully"})
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
