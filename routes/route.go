package routes

import (
	"net/http"

	"github.com/couchbase-examples/golang-quickstart/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controllers struct {
	AirlineController *controllers.AirlineController
	AirportController *controllers.AirportController
	RouteController   *controllers.RouteController
}

func SetupCollectionRoutes(router *gin.Engine, controllers Controllers) {
	// Apply CORS middleware for cross-origin requests
	router.Use(cors.Default())

	// Airline collection endpoints
	// Insert
	router.POST("api/v1/airline/:id", controllers.AirlineController.InsertDocumentForAirline())
	// Get
	router.GET("api/v1/airline/:id", controllers.AirlineController.GetDocumentForAirline())
	// Update
	router.PUT("api/v1/airline/:id", controllers.AirlineController.UpdateDocumentForAirline())
	// Delete
	router.DELETE("api/v1/airline/:id", controllers.AirlineController.DeleteDocumentForAirline())
	// Get
	router.GET("api/v1/airline/to-airport", controllers.AirlineController.GetAirlinesToAirport())
	// Get
	router.GET("api/v1/airline/list", controllers.AirlineController.GetAirlines())

	// Route collection endpoints
	// Insert
	router.POST("api/v1/route/:id", controllers.RouteController.InsertDocumentForRoute())
	// Get
	router.GET("api/v1/route/:id", controllers.RouteController.GetDocumentForRoute())
	// Update
	router.PUT("api/v1/route/:id", controllers.RouteController.UpdateDocumentForRoute())
	// Delete
	router.DELETE("api/v1/route/:id", controllers.RouteController.DeleteDocumentForRoute())

	// Airport collection endpoints
	// Insert
	router.POST("api/v1/airport/:id", controllers.AirportController.InsertDocumentForAirport())
	// Get
	router.GET("api/v1/airport/:id", controllers.AirportController.GetDocumentForAirport())
	// Update
	router.PUT("api/v1/airport/:id", controllers.AirportController.UpdateDocumentForAirport())
	// Delete
	router.DELETE("api/v1/airport/:id", controllers.AirportController.DeleteDocumentForAirport())
	// Get
	router.GET("api/v1/airport/list", controllers.AirportController.GetAirports())
	// Get
	router.GET("api/v1/airport/direct-connections", controllers.AirportController.GetDirectConnections())

	// Swagger UI and documentation
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/docs/index.html")
	})
}
