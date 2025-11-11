package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupConfig() (map[string]string, *os.File) {
	sample_configs := map[string]string{
		"HOST":           "localhost",
		"PORT":           "6000",
		"MIGRATIONS_URL": "file://test_migration_url",
		"DSN_URL":        "postgres://test_postgres_dsn",
	}

	test_config_file, err := os.CreateTemp("", "test*.env")

	if err != nil {
		panic(err)
	}

	return sample_configs, test_config_file
}

func TestLoadConfig(t *testing.T) {
	sample_configs, test_config_file := setupConfig()
	defer os.Remove(test_config_file.Name())

	for key, value := range sample_configs {
		line := fmt.Sprintf("%s=%s\n", key, value)
		_, err := test_config_file.WriteString(line)
		assert.NoError(t, err)
	}

	test_config_file.Sync()

	t.Run("Config file loads", func(t *testing.T) {
		fileDir := filepath.Dir(test_config_file.Name())
		filename := filepath.Base(test_config_file.Name())
		fileExt := filepath.Ext(test_config_file.Name())
		configType, _ := strings.CutPrefix(fileExt, ".")

		config, err := LoadConfig(fileDir, filename, configType)
		assert.NoError(t, err)

		assert.Equal(t, "localhost", config.HOST)
		assert.Equal(t, "6000", config.PORT)
		assert.Equal(t, "postgres://test_postgres_dsn", config.DSN_URL)
		assert.Equal(t, "", config.DSN_OPTIONS)
		assert.Equal(t, "file://test_migration_url", config.MIGRATIONS_URL)
	})
}
