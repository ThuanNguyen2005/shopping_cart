package utils

import (
	"github.com/gin-gonic/gin"
)

type APIRespond struct {
	Message string
	Status  int
	Data    any
}

func RespondSuccess(ctx *gin.Context, status int, message string, data ...any) {
	respond := APIRespond{
		Message: message,
		Status:  status,
	}
	if len(data) > 0 && data[0] != nil {
		respond.Data = data[0]
	}
	ctx.JSON(status, respond)
}

func RespondStatus(ctx *gin.Context, status int) {
	ctx.Status(status)
}
