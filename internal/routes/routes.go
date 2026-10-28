package routes

import (
	"23.usersapi/internal/middleware"
	"23.usersapi/internal/utils"
	"github.com/gin-gonic/gin"
)

type Routers interface {
	Register(r *gin.RouterGroup)
}

func RegisterAllRouter(r *gin.Engine, routers ...Routers) {
	httpLogger := utils.NewLoggerWithPathAndLevel("internal/logs/http.log", "info")
	recoveryLogger := utils.NewLoggerWithPathAndLevel("internal/logs/recovery.log", "error")
	ratelimitLogger := utils.NewLoggerWithPathAndLevel("internal/logs/limitRate.log", "warn")
	r.Use(middleware.LimitRequestMiddleware(ratelimitLogger), middleware.LoggingMiddleware(httpLogger), middleware.APIKeyMiddleware(), middleware.RecoveryMiddleware(recoveryLogger))
	api := r.Group("/api/v1")
	for _, router := range routers {
		router.Register(api)
	}
}
