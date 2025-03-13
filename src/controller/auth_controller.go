package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.challenge/models/dto"
	"go.challenge/service"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(ser *service.AuthService) *AuthController {
	return &AuthController{service: ser}
}

func (ctrl *AuthController) Login(ctx *gin.Context) {
	var request dto.AuthRequestDto
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
	jwt, err := ctrl.service.Login(request.Username, request.Password)
	if (err != nil) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
