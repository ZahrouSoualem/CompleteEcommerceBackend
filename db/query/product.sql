-- name: GetProducts :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY proname
LIMIT $1
OFFSET $2;

-- name: DeleteProducts :exec
DELETE FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (
   category_id,
   brand_id,
   market_id,
   proname,
   description,
   imageurl,
   price,
   quantity
   
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET proname = $2
WHERE id = $1
RETURNING *;


  