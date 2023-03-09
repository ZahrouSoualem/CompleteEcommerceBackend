-- name: GetReview :one
SELECT * FROM review
WHERE id = $1 LIMIT 1;

-- name: ListReviews :many
SELECT * FROM review
ORDER BY id
LIMIT $1
OFFSET $2;

/*-- name: DeleteReview :exec
DELETE FROM review
WHERE id = $1;*/

-- name: CreateReview :one
INSERT INTO review (
  user_id,
  product_id,
  rating
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateReview :one
UPDATE review SET rating = $2
WHERE id = $1
RETURNING *;