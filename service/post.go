package service

import (
	"context"
	"errors"
	db "gocard/db"
	sqlc "gocard/db/sqlc"
	"log"

	"github.com/google/uuid"
)

func Postposts(ownerId, topicId uuid.UUID, title, content string) error {
	// TODO: check topicId is exist or not.
	arg := sqlc.PostpostsParams{
		OwnerID:       ownerId,
		TopicID:       topicId,
		Content:       content,
		Title:         title,
		CreatedBy:     ownerId,
		LastUpdatedBy: ownerId,
	}

	err := db.Queries.Postposts(context.Background(), arg)
	if err != nil {
		log.Println("[Error] Postposts error: ", err.Error())
		return errors.New("Postposts error")
	}

	return nil
}
