package main

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/migrate"
	"time"
)

var session *gocql.Session


func runMigrations() {
	fmt.Println("Running Migrations...")

	connect()
	defer session.Close()

	ctx := context.Background()
	if err := migrate.Migrate(ctx, session, "/app/src/migrations"); err != nil {
		panic(err)
	}

	fmt.Println("Migrations done!")
}

func connect() {
	var err error

	cluster := gocql.NewCluster("scylla")
	cluster.Keyspace = "jtalk"
	cluster.Consistency = gocql.LocalOne
	cluster.Timeout = 10 * time.Second
	cluster.NumConns = 1
	session, err = cluster.CreateSession()

	if err != nil {
		// Must be because the db is not ready yet
		time.Sleep(10000 * time.Millisecond)
		connect()
		fmt.Println("Connection not available trying again later!")
	}
}