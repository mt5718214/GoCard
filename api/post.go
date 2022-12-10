package api

import (
	db "gocard/db"
	sqlc "gocard/db/sqlc"
	"log"
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
// @Tags 			posts
// @Accept 		json
// @Produce 	json
// @Param 		request body db.PostpostsParams true "PostpostsParams"
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
		c.JSON(400, gin.H{
			"result": "Params error",
		})
		return
	}

	if requestBody.Content = strings.Trim(requestBody.Content, " "); requestBody.Content == "" {
		c.JSON(400, gin.H{
			"result": "Content required",
		})
		return
	}

	if requestBody.Title = strings.Trim(requestBody.Title, " "); requestBody.Title == "" {
		c.JSON(400, gin.H{
			"result": "Title required",
		})
		return
	}

	userId := c.GetString("userId")
	ownerId, err := uuid.Parse(userId)
	if err != nil {
		log.Println("uuid parse error: ", err.Error())
		c.JSON(400, gin.H{
			"result": "Invalid uuid",
		})
		return
	}

	// TODO: check topicId is exist or not.
	arg := sqlc.PostpostsParams{
		OwnerID:       ownerId,
		TopicID:       requestBody.TopicId,
		Content:       requestBody.Content,
		Title:         requestBody.Title,
		CreatedBy:     ownerId,
		LastUpdatedBy: ownerId,
	}

	err = db.Queries.Postposts(c, arg)
	if err != nil {
		log.Println("[Error] Postposts error: ", err.Error())
		c.JSON(400, gin.H{
			"result": "Postposts error",
		})
		return
	}

	c.JSON(201, nil)
}
