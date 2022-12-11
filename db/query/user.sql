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

-- name: GetUserByEmail :one
SELECT id, name, email, password FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at;