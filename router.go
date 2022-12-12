package main

import (
	controllers "gocard/controllers"
	route "gocard/route"
	service "gocard/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRouter() *gin.Engine {
	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router
	v1 := server.Group("/dev/api/v1")
	{
		v1.POST("/login", controllers.AuthHandler)
		v1.POST("/signup", controllers.RegisterHandler)

		// The following routes will be authenticated
		v1.Use(service.JWTAuthMiddleware())

		// route
		route.FollowshipRouter(v1)
	}

	return server
}
