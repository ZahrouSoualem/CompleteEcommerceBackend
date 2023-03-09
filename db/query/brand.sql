-- name: GetBrand :one
SELECT * FROM brand
WHERE id = $1 LIMIT 1;

-- name: ListBrands :many
SELECT * FROM brand
ORDER BY braname
LIMIT $1
OFFSET $2;

-- name: DeleteBrand :exec
DELETE FROM brand
WHERE id = $1;

-- name: CreateBrand :one
INSERT INTO brand (
  braname,
  imageurl
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateBrand :one
UPDATE brand SET BraName = $2
WHERE id = $1 
RETURNING *;