package service

import (
	"database/sql"
	"gocard/db"
	sqlc "gocard/db/sqlc"
	"gocard/enum"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	arg := sqlc.PostTopicParams{
		TopicName:     req.TopicName,
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}
	if _, err := db.Queries.PostTopic(c, arg); err != nil {
		log.Println("[Error] PostTopics error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "postTopics error",
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

type updateTopicReq struct {
	TopicID   string `uri:"topicID" binding:"uuid"`
	TopicName string `json:"topicName"`
}

// UpdateTopic  godoc
// @Summary     update topic name
// @Schemes
// @Description	update topic name
// @Tags        topic
// @Accept      json
// @Produce     json
// @Param       topicID path string true "topicID(uuid)"
// @Param       request body updateTopicReq true "topicName"
// @Success     204
// @Router      /admin/topics/:topicId [put]
// @Security    BearerAuth
func UpdateTopic(c *gin.Context) {
	var (
		req updateTopicReq
		err error
	)
	if err = c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid uuid",
		})
		return
	}
	if err = c.ShouldBindJSON(&req); req.TopicName == "" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid topicName",
		})
		return
	}

	id, _ := uuid.Parse(req.TopicID)
	arg := sqlc.UpdateTopicParams{
		TopicName: req.TopicName,
		ID:        id,
	}

	_, err = db.Queries.UpdateTopic(c, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "topicId is not exist",
			})
			return
		}
		log.Println("[Error] UpdateTopic error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "updateTopic err",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
