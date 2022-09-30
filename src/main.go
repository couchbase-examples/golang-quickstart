package main

import (
	"src/routes"
	"github.com/gin-gonic/gin"
	_ "src/controllers"
	_ "src/docs"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	_ "src/docs"
)

// @title Go Profile API
// @version 1.0
// @description Couchbase Golang Quickstart using Gin Gonic
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	router := gin.Default()

	routes.Profileroute(router)

	router.Run()
}