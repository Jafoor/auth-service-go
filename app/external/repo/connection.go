package repo

import (
	"auth-service/config"
	"errors"
	"fmt"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DB struct {
	ReadDB  *sqlx.DB
	WriteDB *sqlx.DB
	psql    sq.StatementBuilderType
}

func getConnectionString(dbConf config.DBConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		dbConf.User,
		dbConf.Pass,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
	if !dbConf.EnableSSLMode {
		connectionString += " sslmode=disable"
	}
	return connectionString
}

func connect(dbConf config.DBConfig) (*sqlx.DB, error) {
	dbSource := getConnectionString(dbConf)

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, fmt.Errorf("connection error: %v", err)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConf.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon, nil
}

func ConnectDB(conf *config.Config) (*DB, error) {
	readDB, err := connect(conf.DB.Read)

	if err != nil {
		return nil, err
	}

	slog.Info("Connect to read database")

	writeDB, err := connect(conf.DB.Write)
	if err != nil {
		return nil, err
	}

	slog.Info("Connected to write database")

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &DB{
		ReadDB:  readDB,
		WriteDB: writeDB,
		psql:    psql,
	}, nil
}

func CloseDB(db *DB) error {
	if err := db.ReadDB.Close(); err != nil {
		return err
	}

	slog.Info("Disconnected from read database")

	if err := db.WriteDB.Close(); err != nil {
		return err
	}

	slog.Info("Disconnected from write database")

	return nil
}

func (db *DB) HealthCheck() error {
	if err := db.ReadDB.Ping(); err != nil {
		var error *pq.Error
		if errors.As(err, &error) {
			slog.Error("Read DB health check failed", "error", err)
		}
		return fmt.Errorf("read database health check failed: %v", err)
	}

	slog.Info("Read database is healthy")

	if err := db.WriteDB.Ping(); err != nil {
		// If ping fails, log and return the error
		var error *pq.Error
		if errors.As(err, &error) {
			slog.Error("Write DB health check failed", "error", err)
		}
		return fmt.Errorf("write database health check failed: %v", err)
	}

	slog.Info("Write database is healthy")

	return nil

}
