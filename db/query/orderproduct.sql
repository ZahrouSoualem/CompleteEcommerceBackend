-- name: Getordersproduct :one
SELECT * FROM ordersproduct
WHERE orders_product_id = $1 LIMIT 1;

-- name: Listordersproducts :many
SELECT product_id, COUNT(product_id) ,SUM(quantity) as total FROM ordersproduct
WHERE orders_id = $1
GROUP BY product_id
ORDER BY total
LIMIT $2
OFFSET $3;

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


-- name: ListJoinOrderProducts :many
SELECT
  ordersproduct.*,orders.*,users.*
FROM
    ordersproduct
    JOIN products ON ordersproduct.product_id = products.id
    JOIN orders ON ordersproduct.orders_id = orders.id
    JOIN users ON orders.user_id = users.id
    where orders.user_id= 1
    order by orders.created_at;

/*-- name: Updateordersproduct :one
UPDATE ordersproduct SET opinion = $2
WHERE order_product_id = $1
RETURNING *;*/