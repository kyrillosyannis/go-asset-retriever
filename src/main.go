package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	database "go.challenge/config"
	"go.challenge/controller"
	"go.challenge/middleware"
	"go.challenge/repository"
	"go.challenge/service"
)

func main() {

	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}
	r := gin.Default()

	chartRepository := repository.NewChartRepository(db)
	chartService := service.NewChartService(chartRepository)
	chartController := controller.NewChartController(chartService)
	favouriteRepository := repository.NewFavouriteRepository(db)
	favouriteService := service.NewFavouriteService(favouriteRepository)
	favouriteController := controller.NewFavouriteController(favouriteService)
	assetRepository := repository.NewAssetRepository(db)
	assetService := service.NewAssetService(assetRepository)
	assetController := controller.NewAssetController(assetService)
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)
	healthController := controller.NewHealthController(db)

	r.GET("/message", chartController.GetMessageHandler)
	r.GET("/charts", chartController.GetAllCharts)
	r.GET("/users/:userId/favourites", favouriteController.GetFavouritesByUser)
	r.POST("/users/:userId/favourites/:assetId", favouriteController.AddToFavourites)
	r.DELETE("/users/:userId/favourites/:favouriteId", favouriteController.RemoveFromFavourites)
	r.POST("/authenticate", authController.Login)
	r.GET("/health", healthController.Health)

	protected := r.Group("/protected")
	protected.Use(middleware.Authenticate)
	protected.PATCH("/assets/:assetId", assetController.UpdateDescription)

	r.Run(":8081") 
}
