-- name: GetOrders :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
WHERE user_id=$1;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (
  user_id
) VALUES (
  $1
)
RETURNING *;

/*-- name: UpdateOrder :one
UPDATE orders SET opinion = $2
WHERE id = $1
RETURNING *;*/