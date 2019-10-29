package main

import (
	"os"
)

var DB_HOST string
var DB_USERNAME string
var DB_PASSWORD string

func readConfigurations() {

	DB_HOST = os.Getenv("DB_HOST")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
}

func getDbHost() string {
	return DB_HOST
}

func getDbUsername() string {
	return DB_USERNAME
}

func getDbPassword() string {
	return DB_PASSWORD
}
