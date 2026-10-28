package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"23.usersapi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RecoveryMiddleware(recoveryLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				stack_at := extractStacklinePanic(stack)
				recoveryLogger.Error().
					Str("path", ctx.Request.URL.Path).
					Str("method", ctx.Request.Method).
					Str("client_IP", ctx.ClientIP()).
					Str("panic", fmt.Sprintf("%v", err)).
					Str("stack_at", stack_at).
					Str("stack", string(stack)).
					Msg("panic occurred.")
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    utils.ERR_CODE_INTERNAL_SERVER,
					"message": "Please try again later.",
				})
			}
		}()
		ctx.Next()
	}
}

func extractStacklinePanic(stack []byte) string {
	lines := bytes.Split(stack, []byte("\n"))
	for _, line := range lines {
		if bytes.Contains(line, []byte(".go")) &&
			!bytes.Contains(line, []byte("/runtime/")) &&
			!bytes.Contains(line, []byte("/debug")) &&
			!bytes.Contains(line, []byte("recovery_middleware.go")) {
			cleanline := strings.TrimSpace(string(line))
			return cleanline
		}
	}
	return ""
}
