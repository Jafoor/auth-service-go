package repo

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func RunMigrations(db sqlx.DB, migrationPath string, direction migrate.MigrationDirection, dryRun bool) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationPath,
	}

	if dryRun {
		plan, _, err := migrate.PlanMigration(db.DB, "postgres", migrations, direction, 0)

		if err != nil {
			slog.Error("Failed to plan migrations", "error", err)
			return fmt.Errorf("failed to plan migrations: %w", err)
		}

		slog.Info("Dry run: planned migrations", "count", len(plan), "direction", direction)

		for _, m := range plan {
			slog.Info("Planned migration", "id", m.Id, "description", m.Migration)
		}
		return nil
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, direction)

	if err != nil {
		slog.Error("Failed to run migrations", "error", err, "direction", direction)
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	slog.Info("Successfully ran migrations", "count", n, "direction", direction)
	return nil
}
