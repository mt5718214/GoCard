package main

import (
	api "gocard/api"
	route "gocard/route"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	server := gin.Default()

	// router
	v1 := server.Group("/dev/api/v1")
	{
		v1.POST("/login", api.AuthHandler)
		v1.POST("/signup", api.RegisterHandler)

		// The following routes will be authenticated
		//v1.Use(api.JWTAuthMiddleware())

		// route
		route.TodoRouter(v1)
	}

	return server
}
