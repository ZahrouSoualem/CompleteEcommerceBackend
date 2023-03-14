-- name: Getordersproduct :one
SELECT * FROM ordersproduct
WHERE orders_product_id = $1 LIMIT 1;

-- name: Listordersproducts :many
SELECT product_id, COUNT(product_id) ,SUM(quantity) as total FROM ordersproduct
GROUP BY product_id
ORDER BY total
LIMIT $1
OFFSET $2;

-- name: Deleteordersproduct :exec
DELETE FROM ordersproduct
WHERE orders_product_id = $1;

-- name: Createordersproduct :one
INSERT INTO ordersproduct (
  orders_id,
  product_id,
  quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

/*-- name: Updateordersproduct :one
UPDATE ordersproduct SET opinion = $2
WHERE order_product_id = $1
RETURNING *;*/