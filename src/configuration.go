package main

import (
	"os"
)

// DbHost Host to connect to DB
var DbHost string

// DbUsername Username to connect to DB
var DbUsername string

// DbPassword Password to connect to DB
var DbPassword string

func readConfigurations() {

	DbHost = os.Getenv("DB_HOST")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
}

func getDbHost() string {
	return DbHost
}

func getDbUsername() string {
	return DbUsername
}

func getDbPassword() string {
	return DbPassword
}
