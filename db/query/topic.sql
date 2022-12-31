-- name: PostTopics :one
INSERT INTO topics (
    topic_name,
    created_by,
    last_updated_by
) VALUES (
    $1,
    $2,
    $3
) RETURNING id;

-- name: UpdateTopic :one
UPDATE topics 
SET topic_name = $1
WHERE id = $2
RETURNING id, topic_name;