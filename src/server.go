package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

var router *gin.Engine

var session *gocql.Session

func connect() {
	var err error

	cluster := gocql.NewCluster("scylla")

	session, err = cluster.CreateSession()

	if err != nil {
		// Must be because the db is not ready yet
		time.Sleep(10000 * time.Millisecond)
		connect()
	}

	fmt.Println("scylla init done!")
}

func setupRouter() {
	router = gin.Default()
	initRoutes()
}

func main() {
	readConfigurations()
	setupRouter()
	router.Run(":7070")
}
