package databaseconfiglib

type DatabaseConfig struct {
	DatabaseHost string
	DatabasePort int64
	DatabaseUser string
	DatabasePass string
	DatabaseName string
	DisableSSL   bool
}
