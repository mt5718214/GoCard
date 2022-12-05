-- name: PostFollowship :one
INSERT INTO followship (
  follower_id,
  topic_id,
  created_by,
  last_updated_by
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

-- name: DeleteFollowship :exec
DELETE FROM followship
WHERE follower_id = $1
AND topic_id = $2;