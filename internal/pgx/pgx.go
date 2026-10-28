package pgx

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type PgxZeroLog struct {
	Logger            *zerolog.Logger
	SlowQueryDuration time.Duration
}

// func (t *PgxZeroLog) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
// 	log.Printf("đây là phần cần xem\n %v \n Kết thúc\n", msg)
// 	sql, _ := data["sql"].(string)
// 	args, _ := data["args"].([]any)
// 	duration, _ := data["time"].(time.Duration)
// 	baseLogger := t.Logger.With().
// 		Dur("duration", duration).
// 		Str("sql", sql).
// 		Any("args", args)
// 	logger := baseLogger.Logger()
// 	if duration > t.SlowQueryDuration && msg == "Query" {
// 		logger.Warn().Str("event", msg).Msg("query is small.")
// 		return
// 	}
// 	// if msg == "Query" {
// 	logger.Info().Str("event", msg).Msg("success query.")
// 	// }

// }
func (t *PgxZeroLog) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	// Tránh nil pointer dereference nếu chưa pass logger
	if t.Logger == nil {
		return
	}

	// 1. Kiểm tra nếu không có SQL (Connect/Acquire/Release) thì bỏ qua ngay
	sql, _ := data["sql"].(string)
	if sql == "" {
		return
	}

	// 2. Lấy args và duration
	args, _ := data["args"].([]any)
	duration, _ := data["time"].(time.Duration)

	// 3. Tạo logger - BẮT BUỘC ĐƯA "sql" VÀO ĐÂY
	baseLogger := t.Logger.With().
		Dur("duration", duration).
		Str("sql", sql).
		Interface("args", args)

	l := baseLogger.Logger()

	// 4. Ghi log khi thực thi Query / Exec / Prepare
	if (msg == "Query" || msg == "Exec" || msg == "Prepare") && duration > t.SlowQueryDuration {
		l.Warn().Str("event", msg).Msg("Slow SQL Query")
		return
	}

	if msg == "Query" || msg == "Exec" || msg == "Prepare" {
		l.Info().Str("event", msg).Msg("Executed SQL")
		return
	}
}
