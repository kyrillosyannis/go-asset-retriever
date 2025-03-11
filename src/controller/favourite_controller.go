package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.challenge/service"
)

type FavouriteController struct {
	service *service.FavouriteService
}

func NewFavouriteController(ser *service.FavouriteService) *FavouriteController {
	return &FavouriteController{service: ser}
}

func (c *FavouriteController) GetFavouritesByUser(ctx *gin.Context) {
	charts, err := c.service.GetFavouritesByUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, charts)
}
