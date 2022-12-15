// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: post.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const postposts = `-- name: Postposts :exec
INSERT INTO posts (
  owner_id,
  topic_id,
  content,
  title,
  created_by,
  last_updated_by
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
`

type PostpostsParams struct {
	OwnerID       uuid.UUID `json:"owner_id"`
	TopicID       uuid.UUID `json:"topic_id"`
	Content       string    `json:"content"`
	Title         string    `json:"title"`
	CreatedBy     uuid.UUID `json:"created_by"`
	LastUpdatedBy uuid.UUID `json:"last_updated_by"`
}

func (q *Queries) Postposts(ctx context.Context, arg PostpostsParams) error {
	_, err := q.db.ExecContext(ctx, postposts,
		arg.OwnerID,
		arg.TopicID,
		arg.Content,
		arg.Title,
		arg.CreatedBy,
		arg.LastUpdatedBy,
	)
	return err
}
