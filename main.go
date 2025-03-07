package main

import (
	"github.com/gin-gonic/gin"
	"go.challenge/service"
	"go.challenge/controller"
)

func main() {
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
