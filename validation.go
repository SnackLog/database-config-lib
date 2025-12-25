package databaseconfiglib

import "errors"

// ValidateConfig validates the given DatabaseConfig
func ValidateConfig(config DatabaseConfig) error {
	if config.DatabaseHost == "" {
		return errors.New("Missing required field: DatabaseHost")
	}
	if config.DatabaseUser == "" {
		return errors.New("Missing required field: DatabaseUser")
	}
	if config.DatabasePass == "" {
		return errors.New("Missing required field: DatabasePass")
	}
	if config.DatabasePort <= 0 || config.DatabasePort > 65535 {
		return errors.New("Invalid value for DatabasePort: must be between 1 and 65535")
	}

	return nil
}
