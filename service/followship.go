package service

import (
	"errors"
	db "gocard/db"
	sqlc "gocard/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostFollowship(c *gin.Context, userId uuid.UUID, topicID uuid.UUID) error {
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

	if _, err := db.Queries.PostFollowship(c, arg); err != nil {
		return errors.New("something went wrong")
	}
	return nil
}

func DeleteFollowship(c *gin.Context, userId uuid.UUID, topicID uuid.UUID) error {
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

	if err := db.Queries.DeleteFollowship(c, arg); err != nil {
		return errors.New("something went wrong")
	}
	return nil
}
