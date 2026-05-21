package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondSuccess(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, gin.H{
		"status": "success",
		"data":   data,
	})
}

func RespondStatus(ctx *gin.Context, status int) {
	ctx.Status(status)
}
