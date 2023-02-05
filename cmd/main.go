package main

import (
	"time"

	"github.com/gin-gonic/gin"
	routeV1 "github.com/oscarllamas6/go-backend-clean-architecture/api/route/v1"
	"github.com/oscarllamas6/go-backend-clean-architecture/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	// Creamos un grupo para poder versionar la api
	routerVersion1 := gin.Group("v1")

	// Llamamos al setup del router
	routeV1.Setup(env, timeout, db, routerVersion1)

	gin.Run(env.ServerAddress)
}
