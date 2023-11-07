package main

import (
    "src/routes"
    "github.com/gin-gonic/gin"
    _ "src/controllers"
    _ "src/docs"
	"src/config"
    "src/db"
    //swaggerFiles "github.com/swaggo/files"
    //ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Golang Quickstart using Gin Gonic
// @version 1.0
// @description Couchbase Golang Quickstart using Gin Gonic. This API provides operations for multiple collections in the database, including CRUD operations and query examples.
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @query.collection.format multi

func main() {
    router := gin.Default()
	// Initialize the cluster
    cluster := db.InitializeCluster()
    // Initialize the shared scope and pass it to the config package
    sharedScope := db.GetScope(cluster)
    config.InitializeSharedScope(sharedScope)
    routes.SetupCollectionRoutes(router)

    router.Run()
}