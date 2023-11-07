package routes

import (
	"net/http"
	"src/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupCollectionRoutes(router *gin.Engine) {
	// Apply CORS middleware for cross-origin requests
	router.Use(cors.Default())

	// Airline collection endpoints
	// Insert
	router.POST("api/v1/airline/:id", controllers.InsertDocumentForAirline())
	// Get
	router.GET("api/v1/airline/:id", controllers.GetDocumentForAirline())
	// Update
	router.PUT("api/v1/airline/:id", controllers.UpdateDocumentForAirline())
	// Delete
	router.DELETE("api/v1/airline/:id", controllers.DeleteDocumentForAirline())
	// Get
	router.GET("api/v1/airline/list", controllers.GetAirlines())
	// Get
	router.GET("api/v1/airline/to-airport", controllers.GetAirlinesToAirport())

	// Route collection endpoints
	// Insert
	router.POST("api/v1/route/:id", controllers.InsertDocumentForRoute())
	// Get
	router.GET("api/v1/route/:id", controllers.GetDocumentForRoute())
	// Update
	router.PUT("api/v1/route/:id", controllers.UpdateDocumentForRoute())
	// Delete
	router.DELETE("api/v1/route/:id", controllers.DeleteDocumentForRoute())

	// Airport collection endpoints
	// Insert
	router.POST("api/v1/airport/:id", controllers.InsertDocumentForAirport())
	// Get
	router.GET("api/v1/airport/:id", controllers.GetDocumentForAirport())
	// Update
	router.PUT("api/v1/airport/:id", controllers.UpdateDocumentForAirport())
	// Delete
	router.DELETE("api/v1/airport/:id", controllers.DeleteDocumentForAirport())
	// Get
	router.GET("api/v1/airport/list", controllers.GetAirports())
	// Get
	router.GET("api/v1/airport/direct-connections", controllers.GetDirectConnections())
	// Swagger UI and documentation
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/docs/index.html")
	})

}
