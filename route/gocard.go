package route

import (
	"gocard/api"

	"github.com/gin-gonic/gin"
)

func GocardRouter(rg *gin.RouterGroup) {
	gocards := rg.Group("/gocards")

	gocards.GET("", api.ListUsers)
	gocards.GET("/:id", api.GetUser)
}
