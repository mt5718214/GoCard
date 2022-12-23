// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: topic.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const postTopics = `-- name: PostTopics :exec
INSERT INTO topics (
    topic_name,
    created_by,
    last_updated_by
) VALUES (
    $1,
    $2,
    $3
)
`

type PostTopicsParams struct {
	TopicName     string    `json:"topic_name"`
	CreatedBy     uuid.UUID `json:"created_by"`
	LastUpdatedBy uuid.UUID `json:"last_updated_by"`
}

func (q *Queries) PostTopics(ctx context.Context, arg PostTopicsParams) error {
	_, err := q.db.ExecContext(ctx, postTopics, arg.TopicName, arg.CreatedBy, arg.LastUpdatedBy)
	return err
}
