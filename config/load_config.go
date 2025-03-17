package config

import (
	"log/slog"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func loadConfig() error {
	exit := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err := godotenv.Load()

	if err != nil {
		slog.Warn(".env not found, that's okay")
	}

	viper.AutomaticEnv()

	config = &Config{
		Mode:        Mode(viper.GetString("MODE")),
		ServiceName: viper.GetString("SERVICE_NAME"),
		HttpPort:    viper.GetInt("HTTP_PORT"),
		DB: DB{
			Read: DBConfig{
				Host:                viper.GetString("READ_DB_HOST"),
				Port:                viper.GetInt("READ_DB_PORT"),
				Name:                viper.GetString("READ_DB_NAME"),
				User:                viper.GetString("READ_DB_USER"),
				Pass:                viper.GetString("READ_DB_PASS"),
				MaxIdleTimeInMinute: viper.GetInt("READ_DB_MAX_IDLE_TIME_IN_MINUTE"),
				EnableSSLMode:       viper.GetBool("READ_DB_ENABLE_SSL_MODE"),
			},
			Write: DBConfig{
				Host:                viper.GetString("WRITE_DB_HOST"),
				Port:                viper.GetInt("WRITE_DB_PORT"),
				Name:                viper.GetString("WRITE_DB_NAME"),
				User:                viper.GetString("WRITE_DB_USER"),
				Pass:                viper.GetString("WRITE_DB_PASS"),
				MaxIdleTimeInMinute: viper.GetInt("WRITE_DB_MAX_IDLE_TIME_IN_MINUTE"),
				EnableSSLMode:       viper.GetBool("WRITE_DB_ENABLE_SSL_MODE"),
			},
		},
		MigrationSourcePath: viper.GetString("MIGRATION_SOURCE_PATH"),
		Redis: Redis{
			Read: []RedisConfig{
				{
					Host:     viper.GetString("REDIS_READ_HOST"),
					Port:     viper.GetString("REDIS_READ_PORT"),
					Password: viper.GetString("REDIS_READ_PASSWORD"),
				},
			},
			Write: RedisConfig{
				Host:     viper.GetString("REDIS_WRITE_HOST"),
				Port:     viper.GetString("REDIS_WRITE_PORT"),
				Password: viper.GetString("REDIS_WRITE_PASSWORD"),
			},
		},
		JWT: Jwt{
			Secret: viper.GetString("JWT_SECRET"),
			ExpIn:  viper.GetInt("JWT_EXP_IN"),
		},
	}

	v := validator.New()

	if err := v.Struct(config); err != nil {
		exit(err)
	}

	return nil

}
