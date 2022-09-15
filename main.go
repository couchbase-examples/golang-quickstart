package main

import (

	"app/routes"
	"github.com/gin-gonic/gin"
	_ "app/docs"
	_ "app/controllers"

	_ "app/docs"
)


func main() {
	router := gin.Default()

	routes.Profileroute(router)

	router.Run()
}