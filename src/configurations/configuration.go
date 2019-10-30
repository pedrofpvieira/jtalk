package configurations

import (
	"os"
)

// DbHost Host to connect to DB
var DbHost string

// DbUsername Username to connect to DB
var DbUsername string

// DbPassword Password to connect to DB
var DbPassword string

// DbSchema reflects the database schema
var DbSchema string

// ReadConfigurations reads all configurations
func ReadConfigurations() {

	DbHost = os.Getenv("DB_HOST")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbSchema = os.Getenv("DB_SCHEMA")
}

// GetDbHost returns the configured db host
func GetDbHost() string {
	return DbHost
}

// GetDbUsername returns the configured db username
func GetDbUsername() string {
	return DbUsername
}

// GetDbPassword returns the configured db password
func GetDbPassword() string {
	return DbPassword
}

// GetDbSchema returns the configured db schema
func GetDbSchema() string {
	return DbSchema
}
