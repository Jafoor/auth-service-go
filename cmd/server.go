package cmd

import (
	"auth-service/app/external/cache"
	"auth-service/app/external/repo"
	"auth-service/config"
	"auth-service/logger"
	"log/slog"
	"sync"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "auth-service server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	slog.Info("starting server")
	conf := config.GetConfig()
	logger.SetupLogger(conf.ServiceName)

	db, err := repo.ConnectDB(conf)
	if err != nil {
		slog.Error("Unable to connect", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	defer repo.CloseDB(db)

	redisClient := cache.InitRedisClient(conf.Redis)
	defer redisClient.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := startRestServer(conf); err != nil {
			slog.Error("REST server failed", slog.String("error", err.Error()))
		}
	}()

	wg.Wait()

	return nil
}
