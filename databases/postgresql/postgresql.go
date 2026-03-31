package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	cfg "github.com/spf13/viper"
)

// NewPostgresqlDB returns a new pgxpool connection pool.
func NewPostgresqlDB() (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.GetString("postgresql.USER"),
		cfg.GetString("postgresql.PASS"),
		cfg.GetString("postgresql.HOST"),
		cfg.GetInt("postgresql.PORT"),
		cfg.GetString("postgresql.DEFAULT_DB"),
	)

	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("postgresql: parse config: %w", err)
	}

	poolCfg.MaxConns = cfg.GetInt32("postgresql.MAX_CONN")
	poolCfg.MaxConnIdleTime = cfg.GetDuration("postgresql.CONN_MAX_IDLE_TIME") * time.Minute
	poolCfg.MaxConnLifetime = cfg.GetDuration("postgresql.CONN_MAX_LIFETIME") * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetDuration("postgresql.CONN_TIMEOUT")*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("postgresql: connect: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgresql: ping: %w", err)
	}

	return pool, nil
}
