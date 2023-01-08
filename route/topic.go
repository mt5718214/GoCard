package route

import (
	"gocard/service"

	"github.com/gin-gonic/gin"
)

func topicRouter(rg *gin.RouterGroup) {
	topics := rg.Group("/topics")

	topics.POST("/", service.PostTopics)
	topics.PUT("/:topicID", service.UpdateTopic)
	topics.DELETE("/:topicID", service.DeleteTopic)
}
