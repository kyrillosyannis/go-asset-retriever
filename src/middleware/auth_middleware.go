package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.challenge/utils"
)

func Authenticate(ctx *gin.Context) {

	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
		ctx.Abort()
		return
	}
	tokenString := header[len("Bearer "):]
	token, err := utils.VerifyJWT(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if (!ok || !token.Valid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	username, exists := claims["username"]
	if (!exists) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	ctx.Set("username", username)
	ctx.Next()
}
