package main

import (
	"jtalk/src/configurations"
	"jtalk/src/migrations"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func setupRouter() {
	router = gin.Default()
	initRoutes()
}

func main() {
	configurations.ReadConfigurations()
	setupRouter()
	migrations.RunMigrations()
	router.Run(":7070")
}
