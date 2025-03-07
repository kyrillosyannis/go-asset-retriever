package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.challenge/controller"
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
	// Initialize Gin
	r := gin.Default()

	// Initialize service and controller
	chartService := &service.ChartService{}
	chartController := controller.NewChartController(chartService)

	// Define endpoint
	r.GET("/message", chartController.GetMessageHandler)

	// Start server
	r.Run(":8081") // Runs on http://localhost:8080
}
