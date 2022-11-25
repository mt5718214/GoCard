// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID      uuid.UUID `json:"id"`
	OwnerID uuid.UUID `json:"owner_id"`
	PostID  uuid.UUID `json:"post_id"`
	// if this field is 0 mean it is the top-level comment
	MainCommentID   uuid.UUID `json:"main_comment_id"`
	Content         string    `json:"content"`
	CreatedBy       uuid.UUID `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	LastUpdatedBy   uuid.UUID `json:"last_updated_by"`
	LastUpdatedDate time.Time `json:"last_updated_date"`
}

type Followship struct {
	ID              uuid.UUID     `json:"id"`
	FollowerID      uuid.NullUUID `json:"follower_id"`
	TopicID         uuid.NullUUID `json:"topic_id"`
	CreatedBy       uuid.UUID     `json:"created_by"`
	CreatedAt       time.Time     `json:"created_at"`
	LastUpdatedBy   uuid.UUID     `json:"last_updated_by"`
	LastUpdatedDate time.Time     `json:"last_updated_date"`
}

type Like struct {
	ID      uuid.UUID     `json:"id"`
	UserID  uuid.NullUUID `json:"user_id"`
	LikedID uuid.NullUUID `json:"liked_id"`
	// 1 is posts, 2 is comment
	Type            sql.NullInt16 `json:"type"`
	CreatedBy       uuid.UUID     `json:"created_by"`
	CreatedAt       time.Time     `json:"created_at"`
	LastUpdatedBy   uuid.UUID     `json:"last_updated_by"`
	LastUpdatedDate time.Time     `json:"last_updated_date"`
}

type Post struct {
	ID              uuid.UUID `json:"id"`
	OwnerID         uuid.UUID `json:"owner_id"`
	TopicID         uuid.UUID `json:"topic_id"`
	Content         string    `json:"content"`
	Title           string    `json:"title"`
	CreatedBy       uuid.UUID `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	LastUpdatedBy   uuid.UUID `json:"last_updated_by"`
	LastUpdatedDate time.Time `json:"last_updated_date"`
}

type Topic struct {
	ID              uuid.UUID `json:"id"`
	TopicName       string    `json:"topic_name"`
	CreatedBy       uuid.UUID `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	LastUpdatedBy   uuid.UUID `json:"last_updated_by"`
	LastUpdatedDate time.Time `json:"last_updated_date"`
}

type User struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	CreatedBy       uuid.UUID `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	LastUpdatedBy   uuid.UUID `json:"last_updated_by"`
	LastUpdatedDate time.Time `json:"last_updated_date"`
}
