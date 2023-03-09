-- name: GetComment :one
SELECT * FROM comment
WHERE id = $1 LIMIT 1;

-- name: ListComments :many
SELECT * FROM comment
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteComment :exec
DELETE FROM Comment
WHERE id = $1;

-- name: CreateComment :one
INSERT INTO comment (
  review_id,
  opinion
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateComment :one
UPDATE comment SET opinion = $2
WHERE id = $1
RETURNING *;