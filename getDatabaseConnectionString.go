package databaseconfiglib

import "fmt"

// GetDatabaseConnectionString constructs and returns a PostgreSQL connection string
// based on the current active database configuration.
func GetDatabaseConnectionString() string {
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", activeConfig.DatabaseUser, activeConfig.DatabasePass, activeConfig.DatabaseHost, activeConfig.DatabasePort, activeConfig.DatabaseName)
	if activeConfig.DisableSSL {
		dbString += "?sslmode=disable"
	}
	return dbString
}
