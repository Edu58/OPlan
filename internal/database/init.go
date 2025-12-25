package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
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

func RunSeeds(db *pgxpool.Pool, logger logger.Logger) {
	accountTypes := []domain.AccountType{
		{
			Name:   "admin",
			Active: true,
		},
		{
			Name:   "user",
			Active: true,
		},
	}

	query := `
		INSERT INTO account_types (name, active)
		VALUES ($1, $2)
		ON CONFLICT (name) DO NOTHING
	`

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var acc domain.AccountType

	for _, a := range accountTypes {
		row := db.QueryRow(ctx, query, a.Name, a.Active)
		if err := row.Scan(&acc.ID, &acc.Name); err != nil {
			if err == pgx.ErrNoRows {
				logger.WithField("Name", a.Name).Info("Skipped Account Type")
			} else {
				log.Panicf("Error running seeds: %v", err)
				logger.WithField("Error", err.Error()).Fatal("Error running seeds")
			}
		}

		logger.WithField("Name", a.Name).Info("Seeded Account Type")
	}
}
