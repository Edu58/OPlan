package database

import (
	"context"
	"errors"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(context context.Context, config *config.Config, logger logger.Logger) (*pgxpool.Pool, error) {
	logger.Info("Initializing Database")

	cfg, err := pgxpool.ParseConfig(config.DSN_URL)

	if err != nil {
		logger.Err(err)
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context, cfg)

	if err != nil {
		logger.Err(err)
		return nil, err
	}

	logger.Info("Running Database migrations")

	if err := runMigrations(config); err != nil {
		logger.Err(err)

		dbPool.Close()

		return dbPool, err
	}

	return dbPool, nil
}

func runMigrations(config *config.Config) error {
	m, err := migrate.New(config.MIGRATIONS_URL, config.DSN_URL)

	if err != nil {
		return err
	}

	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
