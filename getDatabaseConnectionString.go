package databaseconfiglib

import "fmt"

// GetDatabaseConnectionString constructs and returns a PostgreSQL connection string
// based on the current active database configuration.
func GetDatabaseConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", activeConfig.DatabaseUser, activeConfig.DatabasePass, activeConfig.DatabaseHost, activeConfig.DatabasePort, activeConfig.DatabaseName)
}
