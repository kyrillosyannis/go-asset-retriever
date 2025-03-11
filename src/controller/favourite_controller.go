package controller

import (
	"net/http"
	"strconv"

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

func (c *FavouriteController) AddToFavourites(ctx *gin.Context) {
	userIdStr := ctx.Param("userId")
	assetIdStr := ctx.Param("assetId")
	assetId, err := strconv.ParseInt(assetIdStr, 10, 64)
	userId, err2 := strconv.ParseInt(userIdStr, 10, 64)
	if (err != nil || err2 != nil) {
		return //TODO add an informative message about the error
	}
	c.service.AddToFavourites(assetId, userId)
}

func (c *FavouriteController) RemoveFromFavourites(ctx *gin.Context) {
	favIdStr := ctx.Param("favouriteId")
	favId, err := strconv.ParseInt(favIdStr, 10, 64)
	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.service.RemoveFromFavourites(favId)
	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
