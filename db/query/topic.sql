-- name: PostTopics :exec
INSERT INTO topics (
    topic_name,
    created_by,
    last_updated_by
) VALUES (
    $1,
    $2,
    $3
);