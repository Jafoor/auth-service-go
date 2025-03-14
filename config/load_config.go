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
	}

	v := validator.New()

	if err := v.Struct(config); err != nil {
		exit(err)
	}

	return nil

}
