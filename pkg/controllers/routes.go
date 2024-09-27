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
		computerG.GET("/booked", GetBookedComputers)
		computerG.POST("/", superAdminRequired, CreateComputer)
		computerG.GET("/:id", adminOrSuperAdminRequired, GetComputerByID)
		computerG.GET("", adminOrSuperAdminRequired, GetAllComputers)
		computerG.PUT("/:id", superAdminRequired, UpdateComputer)
		computerG.DELETE("/:id", superAdminRequired, DeleteComputer)
	}

	bookingG := apiG.Group("/bookings")
	{
		bookingG.POST("/", CreateBooking)
		bookingG.GET("/:id", adminOrSuperAdminRequired, GetBookingByID)
		bookingG.GET("/", adminOrSuperAdminRequired, GetAllBookings)
		bookingG.PUT("/:id", adminOrSuperAdminRequired, UpdateBooking)
		bookingG.DELETE("/:id", superAdminRequired, DeleteBooking)
	}
	priceListG := apiG.Group("/price-list")
	{
		priceListG.POST("/", superAdminRequired, CreatePriceList)
		priceListG.GET("/", GetAllPriceLists)
		priceListG.GET("/:id", GetPriceListByID)
		priceListG.PUT("/:id", superAdminRequired, UpdatePriceList)
		priceListG.DELETE("/:id", superAdminRequired, DeletePriceList)
	}

	hourlyPackageG := apiG.Group("/hourly-packages")
	{
		hourlyPackageG.POST("/", superAdminRequired, CreateHourlyPackage)
		hourlyPackageG.GET("/", GetAllHourlyPackages)
		hourlyPackageG.GET("/:id", GetHourlyPackageByID)
		hourlyPackageG.PUT("/:id", superAdminRequired, UpdateHourlyPackage)
		hourlyPackageG.DELETE("/:id", superAdminRequired, DeleteHourlyPackage)
	}

	roleG := apiG.Group("/roles")
	{
		roleG.GET("/:id", superAdminRequired, GetRoleByID)
		roleG.POST("/", superAdminRequired, CreateRole)
		roleG.GET("/", superAdminRequired, GetAllRoles)
	}

	usersBalanceG := apiG.Group("/user-balance")
	{
		usersBalanceG.GET("/:username", adminOrSuperAdminRequired, GetUserBalance)
		usersBalanceG.POST("/add-funds", adminOrSuperAdminRequired, TopUpBalanceByUsername)
		usersBalanceG.DELETE("/:username", superAdminRequired, DeleteUserBalance)
		usersBalanceG.POST("/buy/package/:package_id", BuyTimeWithPackage)
		usersBalanceG.POST("/buy/hour/:category_id", BuyTimePerHour)

	}

	// Роут для покупки пакета времени
	apiG.POST("/purchase/hourly-package", PurchaseHourlyPackage)

	// Роут для покупки времени по цене за час
	apiG.POST("/purchase/time", PurchaseTime)

	return router
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
