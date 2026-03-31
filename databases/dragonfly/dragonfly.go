package dragonfly

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	cfg "github.com/spf13/viper"
)

// NewDFClientSingle Returns new DragonFly (using same driver with redis) client ) only for single instance
func NewDFClientSingle(dbNum int) (db *redis.Client, err error) {

	connStr := fmt.Sprintf("%s:%d", cfg.GetString("dragonfly.HOST"), cfg.GetInt("dragonfly.PORT"))

	db = redis.NewClient(&redis.Options{
		Addr:         connStr,
		Username:     cfg.GetString("dragonfly.USER"),
		Password:     cfg.GetString("dragonfly.PASS"),
		MinIdleConns: cfg.GetInt("dragonfly.MIN_IDLE_CONN"),
		PoolSize:     cfg.GetInt("dragonfly.POOL_SIZE"),
		PoolTimeout:  time.Duration(cfg.GetInt("dragonfly.POOL_TIMEOUT")),
		DB:           dbNum,
	})

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetDuration("postgresql.CONN_TIMEOUT")*time.Second)
	defer cancel()

	err = db.Ping(ctx).Err()

	return
}

func NewDFClientCluster() (db *redis.ClusterClient, err error) {

	opt := &redis.ClusterOptions{
		Addrs:        cfg.GetStringSlice("dragonfly.HOSTS"),
		Username:     cfg.GetString("dragonfly.USER"),
		Password:     cfg.GetString("dragonfly.PASS"),
		MinIdleConns: cfg.GetInt("dragonfly.MIN_IDLE_CONN"),
		PoolSize:     cfg.GetInt("dragonfly.POOL_SIZE"),
		PoolTimeout:  time.Duration(cfg.GetInt("dragonfly.POOL_TIMEOUT")),
		MaxRetries:   5,
	}

	db = redis.NewClusterClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetDuration("postgresql.CONN_TIMEOUT")*time.Second)
	defer cancel()

	err = db.Ping(ctx).Err()

	return
}
