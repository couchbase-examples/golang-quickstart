package routes

import (
	"src/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-contrib/cors"
)

func Profileroute(router *gin.Engine) {
	//health check
	router.Use(cors.Default())
	router.GET("api/v1/health", controllers.Healthcheck())
	//insert
	router.POST("api/v1/profile", controllers.InsertProfile())
	//get
	router.GET("api/v1/profile/:id", controllers.GetProfile())
	//update
	router.PUT("api/v1/profile/:id", controllers.UpdateProfile())
	//delete
	router.DELETE("api/v1/profile/:id", controllers.DeleteProfile())
	//swagger UI
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//search
	router.GET("api/v1/profile/profiles", controllers.SearchProfile())

}
