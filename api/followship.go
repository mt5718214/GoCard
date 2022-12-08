package api

import (
	"fmt"
	db "gocard/db"
	sqlc "gocard/db/sqlc"
	"log"

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
			fmt.Println(r)
			c.JSON(400, gin.H{"result": "UUID format error"})
		}
	}()
	topicID := uuid.MustParse(c.Param("topicId"))
	userId, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		log.Fatal(err)
		c.JSON(400, gin.H{"result": "UUID parse error"})
	} else {
		arg := sqlc.PostFollowshipParams{
			FollowerID: uuid.NullUUID{
				UUID:  userId,
				Valid: true,
			},
			TopicID: uuid.NullUUID{
				UUID:  topicID,
				Valid: true,
			},
			CreatedBy:     userId,
			LastUpdatedBy: userId,
		}
		rows, err := db.Queries.PostFollowship(c, arg)

		if err != nil {
			c.JSON(400, gin.H{
				"result": "something went wrong...",
			})
		} else {

			c.JSON(200, gin.H{
				"result": rows,
			})
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
			fmt.Println(r)
			c.JSON(400, gin.H{
				"result": "UUID format error",
			})
		}
	}()
	topicID := uuid.MustParse(c.Param("topicId"))
	userId, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		log.Fatal(err)
		c.JSON(400, gin.H{"result": "UUID parse error"})
	} else {
		arg := sqlc.DeleteFollowshipParams{
			FollowerID: uuid.NullUUID{
				UUID:  userId,
				Valid: true,
			},
			TopicID: uuid.NullUUID{
				UUID:  topicID,
				Valid: true,
			},
		}
		err := db.Queries.DeleteFollowship(c, arg)

		if err != nil {
			c.JSON(400, gin.H{
				"result": "something went wrong...",
			})
		} else {
			c.JSON(204, gin.H{})
		}
	}
}
