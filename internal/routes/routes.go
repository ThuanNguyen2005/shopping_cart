package routes

import (
	"23.usersapi/internal/middleware"
	"23.usersapi/internal/pkg/logger"
	"23.usersapi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Routers interface {
	Register(r *gin.RouterGroup)
}

func RegisterAllRouter(r *gin.Engine, routers ...Routers) {
	httpLogger := NewLoggerWithPathAndLevel("internal/logs/http.log", "info")
	recoveryLogger := NewLoggerWithPathAndLevel("internal/logs/recovery.log", "error")
	r.Use(middleware.LoggingMiddleware(httpLogger), middleware.APIKeyMiddleware(), middleware.RecoveryMiddleware(recoveryLogger), middleware.LimitRequestMiddleware())
	api := r.Group("/api/v1")
	for _, router := range routers {
		router.Register(api)
	}
}

func NewLoggerWithPathAndLevel(path string, level string) *zerolog.Logger {
	loggerConfig := logger.LoggerConfig{
		Filename:   path,
		Level:      level,
		IsDev:      utils.Getenv("ISDEV", "development"),
		MaxSize:    1,    // megabytes
		MaxBackups: 5,    //file backup khi file hiện tại đầy dữ liệu
		MaxAge:     28,   //days(số ngày để xoá file)
		Compress:   true, // disabled by default(có nén không)
		LocalTime:  true,
	}
	return logger.NewLoggerConfig(loggerConfig)
}
