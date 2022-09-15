package routes

import(
	"app/controllers"
	"github.com/gin-gonic/gin"
	//"net/http"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	//"github.com/go-openapi/runtime/middleware"
	"github.com/gin-contrib/cors"
)

func Profileroute(router *gin.Engine) {
	//health check
	router.Use(cors.Default())
	router.GET("api/v1/health",controllers.Healthcheck())
	//insert
	router.POST("api/v1/profile", controllers.InsertProfile())
	//get
	router.GET("api/v1/profile/:id", controllers.GetProfile())
	//update
	router.PUT("api/v1/profile/:id", controllers.UpdateProfile())
	router.DELETE("api/v1/profile/:id", controllers.DeleteProfile())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Docs("/docs",http.FileServer(http.Dir("./")))
	//router.GET("/users", controllers.GetAllUsers())

}