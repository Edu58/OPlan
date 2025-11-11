package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
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

		if err != nil {
			t.Errorf("Error writing to temp config file: %v", err)
		}
	}

	test_config_file.Sync()

	t.Run("Config file loads", func(t *testing.T) {
		fileDir := filepath.Dir(test_config_file.Name())
		filename := filepath.Base(test_config_file.Name())
		fileExt := filepath.Ext(test_config_file.Name())
		configType, _ := strings.CutPrefix(fileExt, ".")

		config, err := LoadConfig(fileDir, filename, configType)

		if err != nil {
			t.Errorf("Error loading config: %v", err)
		}

		for key, value := range sample_configs {
			configValues := reflect.ValueOf(config)
			configValue := configValues.FieldByName(key)

			if value != configValue.String() {
				t.Errorf("Expected \"%s\" got: %v", value, configValue.String())
			}
		}
	})
}
