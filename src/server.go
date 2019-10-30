package main

import (

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func setupRouter() {
	router = gin.Default()
	initRoutes()
}

func main() {
	readConfigurations()
	setupRouter()
	runMigrations()
	router.Run(":7070")
}
