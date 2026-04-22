package clickhouse

import (
	"context"
	"errors"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	cfg "github.com/spf13/viper"
)

// NewClickHouseDB Return new Click House client
func NewClickHouseDB() (db driver.Conn, err error) {

	db, err = clickhouse.Open(&clickhouse.Options{
		Addr: cfg.GetStringSlice("clickhouse.HOSTS"),
		Auth: clickhouse.Auth{
			Database: cfg.GetString("clickhouse.DATABASE"),
			Username: cfg.GetString("clickhouse.USER"),
			Password: cfg.GetString("clickhouse.PASS"),
		},
		DialTimeout:  cfg.GetDuration("clickhouse.DIAL_TIMEOUT"),
		MaxOpenConns: cfg.GetInt("clickhouse.MIN_OPEN_CONN"),
	})

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetDuration("clickhouse.CONN_TIMEOUT"))
	defer cancel()

	if err = db.Ping(ctx); err != nil {
		var exception *clickhouse.Exception
		if errors.As(err, &exception) {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}

	return
}
