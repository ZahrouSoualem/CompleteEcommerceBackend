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
    orders.id as orderid,
    users.id as userid,
    users.username,
    users.email,
    users.phone_number,
    orders.created_at
FROM
     orders 
    JOIN users ON orders.user_id = users.id;

-- name: ListOrderByUserID :many
SELECT
    orders.id as orderid,
    users.id as userid,
    users.username,
    users.email,
    users.phone_number,
    orders.created_at
FROM
     orders 
    JOIN users ON orders.user_id = users.id
    where users.id= $1;


-- name: ProductByOrderID :many
SELECT
    orders_product_id,
    orders_id,
    ordersproduct.quantity,
    product_id,
    products.proname,
    products.price
from
    ordersproduct
    join products on ordersproduct.product_id = products.id
where
    ordersproduct.orders_id = $1;

/*-- name: Updateordersproduct :one
UPDATE ordersproduct SET opinion = $2
WHERE order_product_id = $1
RETURNING *;*/