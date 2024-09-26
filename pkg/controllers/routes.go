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
		userG.GET("/", adminOrSuperAdminRequired, GetAllUsers)
		userG.GET("/:id", adminOrSuperAdminRequired, GetUserByID)
		userG.POST("", adminOrSuperAdminRequired, CreateUser)
		userG.PUT("/:id", adminOrSuperAdminRequired, UpdateUserByID)
		userG.DELETE("/:id", superAdminRequired, DeleteUserByID)
	}

	categoryG := apiG.Group("/categories")
	{
		categoryG.GET("/", GetAllCategories)
		categoryG.GET("/:id", adminOrSuperAdminRequired, GetCategoryByID)
		categoryG.POST("", adminOrSuperAdminRequired, CreateCategory)
		categoryG.PUT("/:id", superAdminRequired, UpdateCategory)
		categoryG.DELETE("/:id", superAdminRequired, DeleteCategory)
	}

	computerG := apiG.Group("/computers")
	{
		computerG.GET("/available", GetAvailableComputers)
		computerG.GET("/booked", GetBookedComputersHandler)
		computerG.POST("/", superAdminRequired, CreateComputer)
		computerG.GET("/:id", adminOrSuperAdminRequired, GetComputerByID)
		computerG.GET("", adminOrSuperAdminRequired, GetAllComputers)
		computerG.PUT("/:id", superAdminRequired, UpdateComputer)
		computerG.DELETE("/:id", superAdminRequired, DeleteComputer)
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
	// Роут для покупки пакета времени
	apiG.POST("/api/purchase/hourly-package", PurchaseHourlyPackage)

	// Роут для покупки времени по цене за час
	apiG.POST("/api/purchase/time", PurchaseTime)

	roleG := apiG.Group("/role")
	{
		roleG.GET("/api/roles/:id", GetRole)
		roleG.POST("/api/roles", CreateRole)
		roleG.GET("/api/roles", GetAllRoles)
	}

	return router
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
