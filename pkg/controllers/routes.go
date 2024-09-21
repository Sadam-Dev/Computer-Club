package controllers

import (
	"ComputerClub/configs"
	_ "ComputerClub/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	apiG := router.Group("/api", checkUserAuthentication)

	userG := router.Group("/users")
	{
		userG.GET("/", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUserByID)
		userG.DELETE("/:id", DeleteUserByID)
	}

	categoryG := router.Group("/categories")
	{
		categoryG.GET("/", GetAllCategories)
		categoryG.GET("/:id", GetCategoryByID)
		categoryG.POST("", CreateCategory)
		categoryG.PUT("/:id", UpdateCategory)
		categoryG.DELETE("/:id", DeleteCategory)
	}

	computerG := router.Group("/computers")
	{
		computerG.GET("/available", GetAvailableComputers)
		computerG.POST("/", CreateComputer)
		computerG.GET("/:id", GetComputerByID)
		computerG.GET("", GetAllComputers)
		computerG.PUT("/:id", UpdateComputer)
		computerG.DELETE("/:id", DeleteComputer)
	}

	bookingG := router.Group("/bookings")
	{
		bookingG.POST("/", CreateBooking)
		bookingG.GET("/:id", GetBookingByID)
		bookingG.GET("/", GetAllBookings)
		bookingG.PUT("/:id", UpdateBooking)
		bookingG.DELETE("/:id", DeleteBooking)
	}
	priceListG := router.Group("/price-list")
	{
		priceListG.POST("/", CreatePriceList)
		priceListG.GET("/", GetAllPriceLists)
		priceListG.GET("/:id", GetPriceListByID)
		priceListG.PUT("/:id", UpdatePriceList)
		priceListG.DELETE("/:id", DeletePriceList)
	}

	hourlyPackageG := router.Group("/hourly-packages")
	{
		hourlyPackageG.POST("/", CreateHourlyPackage)
		hourlyPackageG.GET("/", GetAllHourlyPackages)
		hourlyPackageG.GET("/:id", GetHourlyPackageByID)
		hourlyPackageG.PUT("/:id", UpdateHourlyPackage)
		hourlyPackageG.DELETE("/:id", DeleteHourlyPackage)
	}

	return router
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
