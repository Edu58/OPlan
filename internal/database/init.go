package database

import (
	"context"
	"errors"
	"log"

	"github.com/Edu58/Oplan/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(config *config.Config) (*pgxpool.Pool, error) {
	log.Println("Initializing Database")

	dbPool, err := pgxpool.New(context.Background(), config.DSN_URL)

	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return nil, err
	}

	defer dbPool.Close()

	log.Println("Database initialized succesfully")

	if err := runMigrations(config); err != nil {
		log.Printf("Error running migrations: %v", err)
		return dbPool, err
	}

	return dbPool, nil
}

func runMigrations(config *config.Config) error {
	log.Println("Runing Database migration")

	m, err := migrate.New(config.MIGRATIONS_URL, config.DSN_URL)

	if err != nil {
		return err
	}

	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	log.Println("Database migration successful")

	return nil
}
