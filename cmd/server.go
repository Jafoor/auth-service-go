package cmd

import (
	"auth-service/app/domain/user"
	"auth-service/app/external/cache"
	"auth-service/app/external/repo"
	"auth-service/config"
	"auth-service/logger"
	"log/slog"
	"sync"

	migrate "github.com/rubenv/sql-migrate"
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

	repo.RunMigrations(*db.WriteDB, conf.MigrationSourcePath, migrate.Up, false)

	redisClient := cache.InitRedisClient(conf.Redis)
	defer redisClient.Close()

	userRepo := repo.NewUserRepo(db)

	userService := user.NewService(userRepo, conf.JWT)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := startRestServer(conf, userService); err != nil {
			slog.Error("REST server failed", slog.String("error", err.Error()))
		}
	}()

	wg.Wait()

	return nil
}
