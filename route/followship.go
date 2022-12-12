package route

import (
	"gocard/service"

	"github.com/gin-gonic/gin"
)

func FollowshipRouter(rg *gin.RouterGroup) {
	followship := rg.Group("/followship")

	followship.POST("/:topicId", service.PostFollowship)
	followship.DELETE("/:topicId", service.DeleteFollowship)
}
