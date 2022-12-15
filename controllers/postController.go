package controller

import (
	"gocard/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postPostsBodyParmas struct {
	TopicId uuid.UUID
	Content string
	Title   string
}

// Postposts 	godoc
// @Summary 	create a new post
// @Schemes
// @Description create a new post
// @Tags 		posts
// @Accept 		json
// @Produce 	json
// @Param 		request body postPostsBodyParmas true "postPostsBodyParmas"
// @Success 	201
// @Router 		/posts/ [post]
// @Security 	BearerAuth
func Postposts(c *gin.Context) {
	var (
		requestBody postPostsBodyParmas
		err         error
	)
	if err = c.BindJSON(&requestBody); err != nil {
		log.Println("[Error] Postposts bindJson error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Params error",
		})
		return
	}

	if requestBody.Content = strings.Trim(requestBody.Content, " "); requestBody.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Content required",
		})
		return
	}

	if requestBody.Title = strings.Trim(requestBody.Title, " "); requestBody.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Title required",
		})
		return
	}

	userId := c.GetString("userId")
	ownerId, err := uuid.Parse(userId)
	if err != nil {
		log.Println("uuid parse error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid uuid",
		})
		return
	}

	err = service.Postposts(ownerId, requestBody.TopicId, requestBody.Title, requestBody.Content)
	if err != nil {
		log.Println("[Error] Postposts error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}
