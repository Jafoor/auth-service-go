package cache

import (
	"auth-service/config"
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	writeClient *redis.Client
	readClients []*redis.Client
	mu          sync.RWMutex
}

func getUrl(server, port, password string) string {
	return fmt.Sprintf("redis://%s:%s@%s:%s", password, port, server, port)
}

func InitRedisClient(redisConf config.Redis) *RedisClient {
	var readAddrs []string
	for _, item := range redisConf.Read {
		readAddrs = append(readAddrs, getUrl(item.Host, item.Port, item.Password))
	}

	writeAddr := getUrl(redisConf.Write.Host, redisConf.Write.Port, redisConf.Write.Password)

	return NewRedisClient(writeAddr, readAddrs)
}

func NewRedisClient(writeAddr string, readAdds []string) *RedisClient {
	writeClient := redis.NewClient(&redis.Options{
		Addr:     writeAddr,
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	readClients := make([]*redis.Client, len(readAdds))
	for i, addr := range readAdds {
		readClients[i] = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
			PoolSize: 10,
		})
	}

	return &RedisClient{
		writeClient: writeClient,
		readClients: readClients,
	}
}

func (c *RedisClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.writeClient.Close(); err != nil {
		return fmt.Errorf("failed to close write client: %v", err)
	}

	for _, client := range c.readClients {
		if err := client.Close(); err != nil {
			return fmt.Errorf("failed to close read client: %v", err)
		}
	}

	return nil
}

func (c *RedisClient) HealthCheck(ctx context.Context) (map[string]string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	healthStatus := make(map[string]string)

	err := c.writeClient.Ping(ctx).Err()

	if err != nil {
		healthStatus[c.writeClient.Options().Addr] = "unhealthy"
	} else {
		healthStatus[c.writeClient.Options().Addr] = "healthy"
	}

	for _, client := range c.readClients {
		err := client.Ping(ctx).Err()
		if err != nil {
			healthStatus[client.Options().Addr] = "unhealthy"
		} else {
			healthStatus[client.Options().Addr] = "healthy"
		}
	}

	for addr, status := range healthStatus {
		if status == "unhealthy" {
			return healthStatus, fmt.Errorf("unhealthy Redis instance detected: %s", addr)
		}
	}

	return healthStatus, nil
}
