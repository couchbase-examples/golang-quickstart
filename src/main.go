package main

import (
	"src/config"
	"src/controllers"
	_ "src/controllers"
	"src/db"
	_ "src/docs"
	"src/routes"
	services "src/service"

	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Golang Quickstart using Gin Gonic
// @version 1.0
// @description
// @description This API provides operations for multiple collections in the database, including CRUD operations and query examples.
// @description We have a visual representation of the API documentation using Swagger, which allows you to interact with the API's endpoints directly through the browser. It provides a clear view of the API, including endpoints, HTTP methods, request parameters, and response objects.
// @description Click on an individual endpoint to expand it and see detailed information. This includes the endpoint's description, possible response status codes, and the request parameters it accepts.
// @description Trying Out the API
// @description You can try out an API by clicking on the "Try it out" button next to the endpoints.
// @description - Parameters: If an endpoint requires parameters, Swagger UI provides input boxes for you to fill in. This could include path parameters, query strings, headers, or the body of a POST/PUT request.
// @description - Execution: Once you've inputted all the necessary parameters, you can click the "Execute" button to make a live API call. Swagger UI will send the request to the API and display the response directly in the documentation. This includes the response code, response headers, and response body.
// @description Models
// @description <div style="float: left;">Swagger documents the structure of request and response bodies using models. These models define the expected data structure using JSON schema and are extremely helpful in understanding what data to send and expect.
// @description For details on the API, please check the tutorial on the Couchbase Developer Portal: https://developer.couchbase.com/tutorial-quickstart-golang-gin-gonic
// @query.collection.format multi

func main() {
	router := gin.Default()

	// Initialize the cluster
	cluster := db.InitializeCluster()

	// Initialize the shared scope and pass it to the config package
	sharedScope := db.GetScope(cluster)
	config.InitializeSharedScope(sharedScope)

	// Create service instances
	airlineService := services.NewAirlineService(sharedScope)
	airportService := services.NewAirportService(sharedScope)
	routeService := services.NewRouteService(sharedScope)

	// Create controller instances
	airlineController := controllers.NewAirlineController(airlineService)
	airportController := controllers.NewAirportController(airportService)
	routeController := controllers.NewRouteController(routeService)

	// Create a Controllers struct to hold controller instances
	controllers := routes.Controllers{
		AirlineController: airlineController,
		AirportController: airportController,
		RouteController:   routeController,
	}

	// Setup routes and pass the Controllers struct
	routes.SetupCollectionRoutes(router, controllers)

	router.Run(":8080")
}
