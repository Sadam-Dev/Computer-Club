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

	userG := apiG.Group("/users")
	{
		userG.GET("/", checkUserRole("admin", "super_admin"), GetAllUsers)          // Доступно только для администраторов и выше
		userG.GET("/:id", checkUserRole("admin", "super_admin"), GetUserByID)       // Доступно только для администраторов и выше
		userG.POST("", checkUserRole("user", "admin", "super_admin"), CreateUser)   // Доступно только для администраторов
		userG.PUT("/:id", checkUserRole("admin", "super_admin"), UpdateUserByID)    // Доступно только для администраторов
		userG.DELETE("/:id", checkUserRole("admin", "super_admin"), DeleteUserByID) // Доступно только для администраторов
	}

	categoryG := apiG.Group("/categories")
	{
		categoryG.GET("/", GetAllCategories)
		categoryG.GET("/:id", GetCategoryByID)
		categoryG.POST("", CreateCategory)
		categoryG.PUT("/:id", UpdateCategory)
		categoryG.DELETE("/:id", DeleteCategory)
	}

	computerG := apiG.Group("/computers")
	{
		computerG.GET("/available", checkUserRole("admin", "super_admin"), GetAvailableComputers)
		computerG.POST("/", CreateComputer)
		computerG.GET("/:id", GetComputerByID)
		computerG.GET("", GetAllComputers)
		computerG.PUT("/:id", UpdateComputer)
		computerG.DELETE("/:id", DeleteComputer)
	}

	bookingG := apiG.Group("/bookings")
	{
		bookingG.POST("/", CreateBooking)
		bookingG.GET("/:id", GetBookingByID)
		bookingG.GET("/", GetAllBookings)
		bookingG.PUT("/:id", UpdateBooking)
		bookingG.DELETE("/:id", DeleteBooking)
	}
	priceListG := apiG.Group("/price-list")
	{
		priceListG.POST("/", CreatePriceList)
		priceListG.GET("/", GetAllPriceLists)
		priceListG.GET("/:id", GetPriceListByID)
		priceListG.PUT("/:id", UpdatePriceList)
		priceListG.DELETE("/:id", DeletePriceList)
	}

	hourlyPackageG := apiG.Group("/hourly-packages")
	{
		hourlyPackageG.POST("/", CreateHourlyPackage)
		hourlyPackageG.GET("/", GetAllHourlyPackages)
		hourlyPackageG.GET("/:id", GetHourlyPackageByID)
		hourlyPackageG.PUT("/:id", UpdateHourlyPackage)
		hourlyPackageG.DELETE("/:id", DeleteHourlyPackage)
	}

	// Добавляем новые маршруты для создания и удаления
	apiG.POST("/balance/add-funds", TopUpBalanceByUsername)

	usersBalanceG := apiG.Group("/user-balance")
	{
		usersBalanceG.GET("/:username", GetUserBalance)
	}

	return router
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
