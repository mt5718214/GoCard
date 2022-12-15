package route

import (
	controller "gocard/controllers"

	"github.com/gin-gonic/gin"
)

func PostRouter(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.POST("/", controller.Postposts)
}
