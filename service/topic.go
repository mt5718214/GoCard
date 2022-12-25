package service

import (
	"gocard/db"
	sqlc "gocard/db/sqlc"
	"gocard/enum"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postTopicsReq struct {
	TopicName string `json:"topicName" binding:"required"`
}

// PostTopics   godoc
// @Summary     Create a new topic
// @Schemes
// @Description	Create a new topic
// @Tags        topic
// @Accept      json
// @Produce     json
// @Param       request body postTopicsReq true "topicName"
// @Success     201
// @Router      /admin/topics/ [post]
// @Security    BearerAuth
func PostTopics(c *gin.Context) {
	var (
		req postTopicsReq
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "topicName is required",
		})
		return
	}

	arg := sqlc.PostTopicsParams{
		TopicName:     req.TopicName,
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}
	if err := db.Queries.PostTopics(c, arg); err != nil {
		log.Println("[Error] PostTopics error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "postTopics error",
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
