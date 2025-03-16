package cmd

import (
	"auth-service/app/adapter/rest"
	"auth-service/app/adapter/rest/handlers"
	"auth-service/app/adapter/rest/utils"
	"auth-service/app/domain/user"
	"auth-service/config"
	"log/slog"
	"strconv"
)

func startRestServer(
	conf *config.Config,
	user user.Service,
) error {
	utils.InitValidator()
	slog.Info("Starting REST server", slog.String("port", strconv.Itoa(conf.HttpPort)))

	handler := handlers.NewHandler(conf, user)
	server := rest.NewServer(conf, handler)
	server.Start()
	server.Wg.Wait()

	return nil
}
