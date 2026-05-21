package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerConfig struct {
	Filename   string
	Level      string
	IsDev      string
	MaxSize    int  // megabytes
	MaxBackups int  //file backup khi file hiện tại đầy dữ liệu
	MaxAge     int  //days(số ngày để xoá file)
	Compress   bool // disabled by default(có nén không)
	LocalTime  bool
}

func NewLoggerConfig(loggerConfig LoggerConfig) *zerolog.Logger {
	var write io.Writer
	zerolog.TimeFieldFormat = time.RFC3339
	lvl, err := zerolog.ParseLevel(loggerConfig.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)
	if loggerConfig.IsDev == "development" {
		write = &PrettyJson{writer: os.Stdout}
	} else {
		write = &lumberjack.Logger{
			Filename:   loggerConfig.Filename,
			MaxSize:    loggerConfig.MaxSize,    // megabytes
			MaxBackups: loggerConfig.MaxBackups, //file backup khi file hiện tại đầy dữ liệu
			MaxAge:     loggerConfig.MaxAge,     //days(số ngày để xoá file)
			Compress:   loggerConfig.Compress,   // disabled by default(có nén không)
			LocalTime:  loggerConfig.LocalTime,
		}
	}
	logger := zerolog.New(write).With().Timestamp().Logger()
	return &logger
}

type PrettyJson struct {
	writer io.Writer
}

func (w *PrettyJson) Write(p []byte) (n int, err error) {
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, p, "", "  "); err != nil {
		return w.writer.Write(p)
	}
	return w.writer.Write(prettyJson.Bytes())
}
