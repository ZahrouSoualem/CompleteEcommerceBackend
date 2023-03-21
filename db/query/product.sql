-- name: GetProducts :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListJoinProducts :many
SELECT
    products.id, 
    category_id, 
    brand_id, 
    market_id, 
    proname, 
    description, 
    products.imageurl, 
    price, 
    quantity, 
    catname,
    marketname, 
    email, 
    braname, 
    brand.imageurl 
FROM
    products
    JOIN market ON products.market_id = market.id
    JOIN brand ON products.brand_id = brand.id 
    JOIN category ON products.category_id = category.id
ORDER BY products.id 
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


-- name: UpdateSoldProduct :one
UPDATE products 
SET quantity = quantity + sqlc.arg(amount) 
WHERE id = sqlc.arg(id)
AND quantity + sqlc.arg(amount) >= 0
RETURNING *;

/* UPDATE products 
SET quantity = quantity + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *; */


-- name: UpdateProductName :one
UPDATE products SET proname = $2
WHERE id = $1
RETURNING *;

/*-- name: UpdateProductName :one
UPDATE products SET price = $2
WHERE id = $1
RETURNING *; */


  