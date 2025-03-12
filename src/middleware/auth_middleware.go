package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.challenge/utils"
)

func Authenticate(ctx *gin.Context) {

	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
		ctx.Abort()
		return
	}
	token := header[len("Bearer "):]
	claims, err := utils.VerifyJWT(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	subject, err := claims.Claims.GetSubject()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	fmt.Println(subject)
	ctx.Set("username", subject)
	ctx.Next()
}
