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

type Config struct {
	Mode        Mode
	ServiceName string
	HttpPort    int
}

var config *Config

func GetConfig() *Config {
	configOnce.Do(func() {
		loadConfig()
	})
	return config
}
