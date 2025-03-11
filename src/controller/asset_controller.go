package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.challenge/models/dto"
	"go.challenge/service"
)

type AssetController struct {
	service *service.AssetService
}

func NewAssetController(ser *service.AssetService) *AssetController {
	return &AssetController{service: ser}
}

func (ctrl *AssetController) UpdateDescription(ctx *gin.Context) {
    idStr := ctx.Param("assetId")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
        return
    }
    var request dto.AssetDto
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    asset, err := ctrl.service.UpdateDescription(id, request.Description)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "id":          asset.Id,
        "description": asset.Description,
    })
}
