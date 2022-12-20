package route

import (
	controller "gocard/controllers"

	"github.com/gin-gonic/gin"
)

func followshipRouter(rg *gin.RouterGroup) {
	followship := rg.Group("/followship")

	followship.POST("/:topicId", controller.PostFollowship)
	followship.DELETE("/:topicId", controller.DeleteFollowship)
}
