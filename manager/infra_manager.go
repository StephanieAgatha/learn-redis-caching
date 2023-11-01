package manager

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gookit/slog"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"os"
	"redishop/config"
)

type InfraManager interface {
	Connect() *sql.DB
	RedisClient() *redis.Client
}

type infraManager struct {
	db          *sql.DB
	cfg         *config.Config
	redisClient *redis.Client
}

func (i *infraManager) Connect() *sql.DB {
	return i.db
}

func (i *infraManager) RedisClient() *redis.Client {
	//TODO implement me
	return i.redisClient
}

func (i *infraManager) initdb() error {
	//init dsn disini
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.DbConfig.Host,
		i.cfg.DbConfig.Port,
		i.cfg.DbConfig.Username,
		i.cfg.DbConfig.Password,
		i.cfg.DbConfig.DBName,
	)

	//sql open
	db, err := sql.Open(i.cfg.DBDriver, dsn)
	if err != nil {
		slog.Errorf("Failed to open db %v", err.Error())
		return err
	}
	i.db = db
	return nil
}

func (i *infraManager) initRedis() error {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     i.cfg.RedisConfig.Host,
		Password: i.cfg.RedisConfig.Password,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		slog.Errorf("Failed to connect to Redis: %v", err.Error())
		return err
	}
	slog.Infof("Redis connected to %v", os.Getenv("REDIS_HOST"))

	i.redisClient = client
	return nil
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	if err := conn.initdb(); err != nil {
		return nil, fmt.Errorf("Failed on infra manager %v", err.Error())
	}

	if err := conn.initRedis(); err != nil {
		return nil, fmt.Errorf("Failed to initialize Redis: %v", err.Error())
	}

	return conn, nil
}
