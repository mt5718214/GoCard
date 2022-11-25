package route

import (
	"gocard/api"

	"github.com/gin-gonic/gin"
)

func TodoRouter(rg *gin.RouterGroup) {
	gocards := rg.Group("/gocards")

	gocards.GET("", api.GetTodoLists)
	gocards.GET("/:id", api.GetTodoList)
	gocards.POST("", api.PostTodo)
	gocards.PUT("/:id", api.PutTodo)
	gocards.DELETE("/:id", api.DeleteTodo)
}
