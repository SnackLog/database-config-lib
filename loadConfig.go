package databaseconfiglib

import (
	"fmt"
	"os"
	"strconv"
)

var activeConfig DatabaseConfig

// LoadConfig Load configuration from environment variables and validate it
func LoadConfig() error {
	var loadedConfig DatabaseConfig
	var err error

	loadedConfig.DatabaseHost = os.Getenv("DATABASE_CONFIG_DB_HOST")

	loadedConfig.DatabasePort, err = strconv.ParseInt(os.Getenv("DATABASE_CONFIG_DB_PORT"), 10, 64)
	if err != nil {
		return fmt.Errorf("Invalid port value: %v", err)
	}

	loadedConfig.DatabaseUser = os.Getenv("DATABASE_CONFIG_DB_USER")
	loadedConfig.DatabasePass = os.Getenv("DATABASE_CONFIG_DB_PASS")
	loadedConfig.DatabaseName = os.Getenv("DATABASE_CONFIG_DB_NAME")
	loadedConfig.DisableSSL = os.Getenv("DATABASE_CONFIG_DISABLE_SSL") == "true"

	if loadedConfig.DatabasePort == 0 {
		loadedConfig.DatabasePort = 5432 // default port
	}

	err = ValidateConfig(loadedConfig)
	if err != nil {
		return fmt.Errorf("Config validation failed: %v", err)
	}
	activeConfig = loadedConfig
	return nil
}

// SetConfig sets the active configuration.
// Note: SetConfig does not perform any validation; callers must ensure that
// the provided config has been validated (for example, by calling ValidateConfig)
// before invoking this function.
func SetConfig(config DatabaseConfig) {
	activeConfig = config
}

// GetConfig Return a copy of the active configuration
func GetConfig() DatabaseConfig {
	return activeConfig
}
