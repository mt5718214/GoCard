package controller

import (
	"gocard/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//@followship	/api/v1

// PostFollowship godoc
// @Summary		PostFollowship by toipcId
// @Schemes
// @Description	PostFollowship by toipcId
// @Tags		followship
// @Accept		json
// @Produce		json
// @Param       topicID		path		string		true  "topicID(uuid)"
// @Success	  	200			{string}	json		"{"result":{}}"
// @Router		/followship/:topicId [post]
// @Security 	BearerAuth
func PostFollowship(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "UUID format error"})
		}
	}()
	topicID := uuid.MustParse(c.Param("topicId"))
	if userId, err := uuid.Parse(c.GetString("userId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "UUID parse error"})
	} else {
		if err := service.PostFollowship(c, userId, topicID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusAccepted, gin.H{})
		}
	}

}

// DeleteFollowship godoc
// @Summary	DeleteFollowship by toipcId
// @Schemes
// @Description	DeleteFollowship by toipcId
// @Tags			followship
// @Accept			json
// @Produce		json
// @Param        topicID   path      string  true  "topicID(uuid)"
// @Success	  204
// @Router			/followship/:topicId [delete]
// @Security BearerAuth
func DeleteFollowship(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(400, gin.H{"message": "UUID format error"})
		}
	}()
	topicID := uuid.MustParse(c.Param("topicId"))
	if userId, err := uuid.Parse(c.GetString("userId")); err != nil {
		log.Fatal(err)
		c.JSON(400, gin.H{"message": "UUID parse error"})
	} else {
		if err := service.DeleteFollowship(c, userId, topicID); err != nil {
			c.JSON(400, gin.H{"message": "something went wrong..."})
		} else {
			c.JSON(204, gin.H{})
		}
	}
}
