-- name: GetMarket :one
SELECT * FROM market
WHERE id = $1 LIMIT 1;

-- name: ListMarkets :many
SELECT * FROM market
ORDER BY marketname
LIMIT $1
OFFSET $2;

-- name: DeleteMarket :exec
DELETE FROM market
WHERE id = $1;

-- name: CreateMarket :one
INSERT INTO market (
  marketname,
   email,
   password
   
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMarket :one
UPDATE market SET marketname = $2
WHERE id = $1
RETURNING *;


  