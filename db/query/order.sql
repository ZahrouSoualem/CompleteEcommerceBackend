-- name: GetOrder :one
SELECT * FROM "order"
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM "order"
WHERE user_id=$1
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteOrder :exec
DELETE FROM "order"
WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO "order" (
  user_id
) VALUES (
  $1
)
RETURNING *;

/*-- name: UpdateOrder :one
UPDATE "order" SET opinion = $2
WHERE id = $1
RETURNING *;*/