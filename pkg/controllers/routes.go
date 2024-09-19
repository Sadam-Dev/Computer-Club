package controllers

import (
	"ComputerClub/configs"
	_ "ComputerClub/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
		computerG.GET("/booking", GetBookingComputers)
		computerG.POST("", CreateComputer)
	}

	bookingG := router.Group("/bookings")
	{
		bookingG.POST("", CreateBooking)
	}

	return router
}

func WelcomeToTheClub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to the club, buddy",
	})
}
