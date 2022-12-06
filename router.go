package main

import (
	api "gocard/api"
	docs "gocard/docs"
	route "gocard/route"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRouter() *gin.Engine {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "gocard API doc"
	docs.SwaggerInfo.Description = "This is goCard. You can visit the GitHub repository at https://github.com/mt5718214/GoCard"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router
	v1 := server.Group("/dev/api/v1")
	{
		v1.POST("/login", api.AuthHandler)
		v1.POST("/signup", api.RegisterHandler)

		// The following routes will be authenticated
		v1.Use(api.JWTAuthMiddleware())

		// route
		route.FollowshipRouter(v1)
	}

	return server
}
