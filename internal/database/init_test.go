package database

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestRunMigrationsInvalidConfig(t *testing.T) {

	config := config.Config{
		MIGRATIONS_URL: "invalid migrations filepath",
	}

	err := runMigrations(&config)
	assert.Error(t, err)
}

func TestInitDBInvalidConfig(t *testing.T) {
	var buf bytes.Buffer
	test_logger := logger.NewLogger(&buf)

	config := config.Config{
		DSN_URL: "invalid dsn",
	}

	pool, err := InitDB(context.Background(), &config, test_logger)
	assert.Nil(t, pool)
	assert.Error(t, err)
	assert.NotEmpty(t, buf)
}

func TestInitDB(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx, "postgres:latest",
		postgres.WithDatabase("oplan_dev"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithSQLDriver("pgx"),
		postgres.BasicWaitStrategies(),
	)

	require.NoError(t, err)

	defer pgContainer.Terminate(ctx)

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")

	if err != nil {
		require.NoError(t, err)
	}

	tmpMigDir := os.TempDir()

	defer os.Remove(tmpMigDir)

	upSQL := `CREATE TABLE IF NOT EXISTS users(
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(255) NOT NULL,
	active BOOLEAN NOT NULL
	);`

	if err := os.WriteFile(filepath.Join(tmpMigDir, "0001_create_users.up.sql"), []byte(upSQL), 0o644); err != nil {
		require.NoError(t, err)
	}

	// also provide a down to be nice (optional)
	downSQL := `DROP TABLE IF EXISTS users;`
	if err := os.WriteFile(filepath.Join(tmpMigDir, "0001_create_account_types.down.sql"), []byte(downSQL), 0o644); err != nil {
		t.Fatalf("write down migration: %v", err)
	}

	fmt.Println(dsn)

	config := config.Config{
		DSN_URL:        dsn,
		MIGRATIONS_URL: "file://" + tmpMigDir,
		LOGGER_LEVEL:   "info",
	}

	var buf bytes.Buffer
	logger := logger.NewLoggerWithLevel(config.LOGGER_LEVEL, &buf)

	pgxPool, err := InitDB(ctx, &config, logger)

	assert.NoError(t, err)
	assert.NotNil(t, pgxPool)

	defer pgxPool.Close()

	if err := pgxPool.Ping(ctx); err != nil {
		assert.NoError(t, err)
	}

	var tableExists bool
	conn, err := pgxPool.Acquire(ctx)
	defer conn.Release()

	if err := pgxPool.Ping(ctx); err != nil {
		assert.NoError(t, err)
	}

	rows := conn.QueryRow(ctx, `SELECT EXISTS (
	    SELECT FROM information_schema.tables
	    WHERE  table_schema = 'public'
	    AND    table_name   = 'users'
	);`)

	if err := rows.Scan(&tableExists); err != nil {
		assert.NoError(t, err)
	}

	assert.True(t, tableExists)
}
