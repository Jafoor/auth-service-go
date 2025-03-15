package config

import "sync"

var configOnce = sync.Once{}

type Mode string

const (
	Dev         Mode = "dev"
	Prod        Mode = "prod"
	DebugMode   Mode = "debug"
	ReleaseMode Mode = "release"
)

type RedisConfig struct {
	Host     string `json:"host" validate:"required"`
	Port     string `json:"port" validate:"required"`
	Password string `json:"password" `
}

type Redis struct {
	Read  []RedisConfig `json:"read" validate:"required"`
	Write RedisConfig   `json:"write" validate:"required"`
}

type Config struct {
	Mode        Mode
	ServiceName string
	HttpPort    int
	Redis       Redis
}

var config *Config

func GetConfig() *Config {
	configOnce.Do(func() {
		loadConfig()
	})
	return config
}
