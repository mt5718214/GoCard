package route

import (
	"gocard/api"

	"github.com/gin-gonic/gin"
)

func PostRouter(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.POST("/", api.Postposts)
}
