package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.challenge/service"
)

// MessageController handles HTTP requests.
type ChartController struct {
	service *service.ChartService
}

// NewMessageController creates a new controller instance.
func NewChartController(svc *service.ChartService) *ChartController {
	return &ChartController{service: svc}
}

// GetMessageHandler handles the GET request to return a message.
func (c *ChartController) GetMessageHandler(ctx *gin.Context) {
	message := c.service.GetMessage()
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
