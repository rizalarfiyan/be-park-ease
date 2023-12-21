package database

import (
	"be-park-ease/config"
	"be-park-ease/logger"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var postgresConn *pgxpool.Pool
var postgresLog zerolog.Logger

func InitPostgres(ctx context.Context) {
	postgresLog = logger.Get("postgres")
	postgresLog.Info().Msg("Connect Postgres server...")
	conf := config.Get()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Name)
	connConf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		postgresLog.Fatal().Err(err).Msg("Postgres configration problem")
	}

	connConf.MaxConnIdleTime = conf.DB.ConnectionIdle
	connConf.MaxConnLifetime = conf.DB.ConnectionLifetime
	connConf.MaxConns = int32(conf.DB.MaxOpen)
	connConf.MinConns = int32(conf.DB.MaxIdle)

	db, err := pgxpool.NewWithConfig(ctx, connConf)
	if err != nil {
		postgresLog.Fatal().Err(err).Msg("Postgres connection problem")
	}

	postgresConn = new(pgxpool.Pool)
	postgresConn = db

	postgresLog.Info().Msg("Postgres connected...")
}

func GetPostgres() *pgxpool.Pool {
	return postgresConn
}

func PostgresIsConnected() bool {
	ctx := context.Background()
	err := postgresConn.Ping(ctx)
	if err != nil {
		postgresLog.Error().Err(err).Msg("Postgres fails health check")
		return false
	}
	return true
}
