package db

import (
	"context"
	"fmt"
	"time"

	"23.usersapi/internal/config"
	"23.usersapi/internal/db/sqlc"
	"23.usersapi/internal/pgx"
	"23.usersapi/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

var DB sqlc.Querier
var DBpool *pgxpool.Pool

func InitDB() error {
	connStr := config.NewConfig().DNS()
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("Error parse DB config: %w", err)
	}
	sqlQueryLogger := utils.NewLoggerWithPathAndLevel("internal/logs/sqlQuery.log", "info")
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: &pgx.PgxZeroLog{
			Logger:            sqlQueryLogger,
			SlowQueryDuration: 500 * time.Millisecond,
		},
		LogLevel: tracelog.LogLevelDebug, // BẮT BUỘC BẬT DEBUG Ở ĐÂY
	}
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConns = 50
	config.MinConns = 10
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	DBpool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("Error create pool db: %w", err)
	}
	if err := DBpool.Ping(ctx); err != nil {
		return fmt.Errorf("Error ping db: %w", err)
	}
	DB = sqlc.New(DBpool)

	return nil
}
