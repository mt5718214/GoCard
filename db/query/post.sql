-- name: Postposts :exec
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
);