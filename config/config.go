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

type DBConfig struct {
	Host                string `json:"host"                    validate:"required"`
	Port                int    `json:"port"                    validate:"required"`
	Name                string `json:"name"                    validate:"required"`
	User                string `json:"user"                    validate:"required"`
	Pass                string `json:"pass"                    validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

type DB struct {
	Read  DBConfig `json:"read"  validate:"required"`
	Write DBConfig `json:"write" validate:"required"`
}

type RedisConfig struct {
	Host     string `json:"host" validate:"required"`
	Port     string `json:"port" validate:"required"`
	Password string `json:"password" `
}

type Redis struct {
	Read  []RedisConfig `json:"read" validate:"required"`
	Write RedisConfig   `json:"write" validate:"required"`
}

type Jwt struct {
	Secret string `json:"secret" validate:"required"`
	ExpIn  int    `json:"exp_in" validate:"required"`
}

type Config struct {
	Mode                Mode   `json:"mode"`
	ServiceName         string `json:"service_name"`
	HttpPort            int    `json:"http_port" `
	DB                  DB     `json:"db"  validate:"required"`
	MigrationSourcePath string `json:"migration_source_path" validate:"required"`
	Redis               Redis  `json:"redis"                 validate:"required"`
	JWT                 Jwt    `json:"jwt"                   validate:"required"`
}

var config *Config

func GetConfig() *Config {
	configOnce.Do(func() {
		loadConfig()
	})
	return config
}
