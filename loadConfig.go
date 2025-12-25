package databaseconfiglib

import (
	"fmt"
	"os"
)

var activeConfig DatabaseConfig

// LoadConfig Load configuration from environment variables and validate it
func LoadConfig() error {
	var loadedConfig DatabaseConfig
	loadedConfig.DatabaseHost = os.Getenv("DATABASE_CONFIG_DB_HOST")
	loadedConfig.DatabasePort = os.Getenv("DATABASE_CONFIG_DB_PORT")
	loadedConfig.DatabaseUser = os.Getenv("DATABASE_CONFIG_DB_USER")
	loadedConfig.DatabasePass = os.Getenv("DATABASE_CONFIG_DB_PASS")

	if loadedConfig.DatabasePort == "" {
		loadedConfig.DatabasePort = "5432" // default port
	}

	err := ValidateConfig(loadedConfig)
	if err != nil {
		return fmt.Errorf("Config validation failed: %v", err)
	}
	activeConfig = loadedConfig
	return nil
}

// SetConfig Set the active configuration
func SetConfig(config DatabaseConfig) {
	activeConfig = config
}

// GetConfig Return a copy of the active configuration
func GetConfig() DatabaseConfig {
	return activeConfig
}
