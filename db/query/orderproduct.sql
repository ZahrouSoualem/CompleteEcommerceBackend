-- name: GetOrderProduct :one
SELECT * FROM OrderProduct
WHERE order_product_id = $1 LIMIT 1;

-- name: ListOrderProducts :many
SELECT product_id, COUNT(product_id) ,SUM(quantity) as total FROM OrderProduct
GROUP BY product_id
ORDER BY total
LIMIT $1
OFFSET $2;

-- name: DeleteOrderProduct :exec
DELETE FROM OrderProduct
WHERE order_product_id = $1;

-- name: CreateOrderProduct :one
INSERT INTO OrderProduct (
  order_id,
  product_id,
  quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

/*-- name: UpdateOrderProduct :one
UPDATE OrderProduct SET opinion = $2
WHERE order_product_id = $1
RETURNING *;*/