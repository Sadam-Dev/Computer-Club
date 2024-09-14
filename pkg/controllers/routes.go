package controllers

import (
	"ComputerClub/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunRoutes() error {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/welcome", WelcomeToTheClub)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := router.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUserByID)
		userG.DELETE("/:id", DeleteUserByID)
	}

	computerG := router.Group("/computers")
	{
		computerG.GET("/available", GetAvailableComputers)
		computerG.POST("", CreateComputer)
	}

	bookingG := router.Group("/bookings")
	{
		bookingG.POST("", CreateBooking)
	}

	err := router.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))

	if err != nil {
		return err
	}

	return nil
}

func WelcomeToTheClub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to the club, buddy",
	})
}
