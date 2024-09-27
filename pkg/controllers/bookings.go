package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateBooking godoc
// @Summary      Create a new booking
// @Description  Creates a new booking for the logged-in user.
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param        booking  body      models.SwagBooking  true  "Booking Details"
// @Success      201      {object}  models.Booking
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /bookings [post]
func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.BindJSON(&booking); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrUnauthorized)
		return
	}

	booking.UserID = userID

	if err := service.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Бронирование создано успешно", "booking": booking})
}

// GetBookingByID godoc
// @Summary      Get booking by ID
// @Description  Retrieve a booking by its ID.
// @Tags         bookings
// @Produce      json
// @Param        id   path      int  true  "Booking ID"
// @Success      200  {object}  models.Booking
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /bookings/{id} [get]
func GetBookingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	booking, err := service.GetBookingByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, booking)
}

// GetAllBookings godoc
// @Summary      Get all bookings
// @Description  Retrieve a list of all bookings.
// @Tags         bookings
// @Produce      json
// @Success      200  {array}   models.Booking
// @Failure      500  {object}  map[string]string
// @Router       /bookings [get]
func GetAllBookings(c *gin.Context) {
	bookings, err := service.GetAllBookings()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// UpdateBooking godoc
// @Summary      Update a booking
// @Description  Update an existing booking by ID.
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Booking ID"
// @Param        booking  body      models.SwagBooking  true  "Updated Booking Details"
// @Success      200      {object}  models.Booking
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /bookings/{id} [put]
func UpdateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.BindJSON(&booking); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	booking.ID = uint(id)

	if err := service.UpdateBooking(booking); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Бронирование обновлено успешно", "booking": booking})
}

// DeleteBooking godoc
// @Summary      Delete a booking
// @Description  Delete a booking by its ID.
// @Tags         bookings
// @Produce      json
// @Param        id   path      int  true  "Booking ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /bookings/{id} [delete]
func DeleteBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if err := service.DeleteBooking(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Бронирование удалено успешно"})
}
