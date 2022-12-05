package route

import (
	"gocard/api"

	"github.com/gin-gonic/gin"
)

func FollowshipRouter(rg *gin.RouterGroup) {
	gocards := rg.Group("/followship")

	gocards.POST("/:topicId", api.PostFollowship)
	gocards.DELETE("/:topicId", api.DeleteFollowship)
}
