package databaseconfiglib

import "errors"

// ValidateConfig Validate the given ServiceConfig
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

	return nil
}
