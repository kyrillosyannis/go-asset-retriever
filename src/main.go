package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	database "go.challenge/config"
	"go.challenge/controller"
	"go.challenge/repository"
	"go.challenge/service"
)

func main() {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}
	// Initialize Gin
	r := gin.Default()

	// Initialize service and controller
	chartRepository := repository.NewChartRepository(db)
	chartService := service.NewChartService(chartRepository)
	chartController := controller.NewChartController(chartService)
	favouriteRepository := repository.NewFavouriteRepository(db)
	favouriteService := service.NewFavouriteService(favouriteRepository)
	favouriteController := controller.NewFavouriteController(favouriteService)
	assetRepository := repository.NewAssetRepository(db)
	assetService := service.NewAssetService(assetRepository)
	assetController := controller.NewAssetController(assetService)

	// Define endpoint
	r.GET("/message", chartController.GetMessageHandler)
	r.GET("/charts", chartController.GetAllCharts)
	r.GET("/users/:userId/favourites", favouriteController.GetFavouritesByUser)
	r.POST("/users/:userId/favourites/:assetId", favouriteController.AddToFavourites)
	r.DELETE("/users/:userId/favourites/:favouriteId", favouriteController.RemoveFromFavourites)
	r.PATCH("/assets/:assetId", assetController.UpdateDescription)

	// Start server
	r.Run(":8081") // Runs on http://localhost:8081
}
