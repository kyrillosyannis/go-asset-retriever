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

	r.GET("/message", chartController.GetMessageHandler)
	r.GET("/charts", chartController.GetAllCharts)
	r.GET("/users/:userId/favourites", favouriteController.GetFavouritesByUser)
	r.POST("/users/:userId/favourites/:assetId", favouriteController.AddToFavourites)
	r.DELETE("/users/:userId/favourites/:favouriteId", favouriteController.RemoveFromFavourites)
	r.PATCH("/assets/:assetId", assetController.UpdateDescription)

	r.GET("/health", func(c *gin.Context) {
    sqlDB, _ := db.DB()
    err := sqlDB.Ping()
    if err != nil {
        c.JSON(500, gin.H{"status": "database_down"})
        return
    }
    c.JSON(200, gin.H{"status": "up"})
	})

	r.Run(":8081") 
}
