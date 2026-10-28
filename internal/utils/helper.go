package utils

import (
	"os"

	"23.usersapi/internal/pkg/logger"
	"github.com/rs/zerolog"
)

func Getenv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func NewLoggerWithPathAndLevel(path string, level string) *zerolog.Logger {
	loggerConfig := logger.LoggerConfig{
		Filename:   path,
		Level:      level,
		IsDev:      Getenv("ISDEV", ""),
		MaxSize:    1,    // megabytes
		MaxBackups: 5,    //file backup khi file hiện tại đầy dữ liệu
		MaxAge:     28,   //days(số ngày để xoá file)
		Compress:   true, // disabled by default(có nén không)
		LocalTime:  true,
	}
	return logger.NewLoggerConfig(loggerConfig)
}
