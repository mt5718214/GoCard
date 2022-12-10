-- name: PostUser :one
INSERT INTO users (
  name,
  email,
  password,
  created_by,
  last_updated_by
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at;