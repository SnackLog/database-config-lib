package databaseconfiglib

import "fmt"

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", activeConfig.DatabaseUser, activeConfig.DatabasePass, activeConfig.DatabaseHost, activeConfig.DatabasePort, activeConfig.DatabaseName)
}
