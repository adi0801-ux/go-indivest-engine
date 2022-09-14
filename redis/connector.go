package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type ConnectionConfig struct {
	Address  string
	Port     string
	Password string
	Username string
	DBName   int
}

func ConnectToRedis(config *ConnectionConfig) (*Client, error) {
	options := redis.Options{
		Addr:     config.Address,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DBName,
	}

	client := redis.NewClient(&options)

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	connection := Client{store: client}

	return &connection, nil

}

type Client struct {
	store *redis.Client
}
