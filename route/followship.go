package route

import (
	"gocard/service"

	"github.com/gin-gonic/gin"
)

func FollowshipRouter(rg *gin.RouterGroup) {
	gocards := rg.Group("/followship")

	gocards.POST("/:topicId", service.PostFollowship)
	gocards.DELETE("/:topicId", service.DeleteFollowship)
}
