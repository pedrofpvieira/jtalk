package db

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// Session to access the DB
var Session *gocql.Session

// Connect establishes a connection to the DB and returns the session
func Connect() {
	var err error

	// TODO Understand why it does not work reading host and keyspace from configuration.env
	cluster := gocql.NewCluster("scylla")
	cluster.Keyspace = "jtalk"
	cluster.Consistency = gocql.LocalOne
	cluster.Timeout = 10 * time.Second
	cluster.NumConns = 1
	Session, err = cluster.CreateSession()

	if err != nil {
		// Must be because the db is not ready yet
		time.Sleep(10000 * time.Millisecond)
		Connect()
		fmt.Println("Connection ScyllaDB is not available, trying again later!")
	}
}
