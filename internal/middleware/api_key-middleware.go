package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	ENVKey := os.Getenv("API_Key")
	if ENVKey == "" {
		ENVKey = "secret key"
	}
	return func(ctx *gin.Context) {
		APIKey := ctx.GetHeader("X-API-Key")
		if APIKey == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid key",
			})
			return
		}
		if APIKey != ENVKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid key",
			})
			return
		}
		if APIKey == ENVKey {
			ctx.Next()
		}
	}
}
