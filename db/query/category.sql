-- name: GetCategory :one
SELECT * FROM category
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM category
ORDER BY CatName
LIMIT $1
OFFSET $2;

-- name: DeleteCategories :exec
DELETE FROM category
WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO category (
  catname
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateCategory :one
UPDATE category SET CatName = $2
WHERE id = $1 
RETURNING *;