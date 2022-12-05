package api

import (
	"fmt"
	db "gocard/db"
	sqlc "gocard/db/sqlc"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostFollowship(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(400, gin.H{"result": "UUID format error"})
		}
	}()
	//inputUUid, err := uuid.Parse("7b425632-aadb-4aec-8dfc-feca9d978ad2")
	//token, _ := createJWT("sub", inputUUid, "albert")
	//log.Fatal(token)
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
