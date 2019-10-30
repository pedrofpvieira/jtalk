package migrations

import (
	"context"
	"fmt"
	"jtalk/src/db"
	"log"
	"os"

	"github.com/scylladb/gocqlx/migrate"
)

// RunMigrations runs all database migrations if needed
func RunMigrations() {
	db.Connect()
	ctx := context.Background()

	fmt.Println("Running Migrations...")
	if err := migrate.Migrate(ctx, db.Session, getDir()); err != nil {
		panic(err)
	}

	fmt.Println("Migrations done!")
}

func getDir() string {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return path + "/src/migrations"
}
