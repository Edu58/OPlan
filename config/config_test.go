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
	sampleConfigs := map[string]string{
		"HOST":           "localhost",
		"PORT":           "6000",
		"MIGRATIONS_URL": "file://test_migration_url",
		"DSN_URL":        "postgres://test_postgres_dsn",
	}

	testConfigFile, err := os.CreateTemp("", "test*.env")

	if err != nil {
		panic(err)
	}

	return sampleConfigs, testConfigFile
}

func TestLoadConfig(t *testing.T) {
	sampleConfigs, testConfigFile := setupConfig()
	defer os.Remove(testConfigFile.Name())

	for key, value := range sampleConfigs {
		line := fmt.Sprintf("%s=%s\n", key, value)
		_, err := testConfigFile.WriteString(line)
		assert.NoError(t, err)
	}

	err := testConfigFile.Sync()
	if err != nil {
		return
	}

	t.Run("Config file loas", func(t *testing.T) {
		fileDir := filepath.Dir(testConfigFile.Name())
		filename := filepath.Base(testConfigFile.Name())
		fileExt := filepath.Ext(testConfigFile.Name())
		configType, _ := strings.CutPrefix(fileExt, ".")

		config, err := LoadConfig(fileDir, filename, configType)
		assert.NoError(t, err)

		assert.Equal(t, "localhost", config.HOST)
		assert.Equal(t, "6000", config.PORT)
		assert.Equal(t, "postgres://test_postgres_dsn", config.DsnUrl)
		assert.Equal(t, "", config.DsnOptions)
		assert.Equal(t, "file://test_migration_url", config.MigrationsUrl)
	})
}
