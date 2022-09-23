package routes

import(
	"app/controllers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"github.com/gin-contrib/cors"
)

func Profileroute(router *gin.Engine) {
	
	//CORS 
	router.Use(cors.Default())
	//health check
	router.GET("api/v1/health",controllers.Healthcheck())
	//insert
	router.POST("api/v1/profile", controllers.InsertProfile())
	//get
	router.GET("api/v1/profile/:id", controllers.GetProfile())
	//update
	router.PUT("api/v1/profile/:id", controllers.UpdateProfile())
	//delete
	router.DELETE("api/v1/profile/:id", controllers.DeleteProfile())
	//serve swagger UI
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


}