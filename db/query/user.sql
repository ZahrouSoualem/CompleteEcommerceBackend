-- name: GetUsers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
  username,
   email,
   password,
   address,
   city,
   state,
   country,
   zip_code,
   phone_number
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET username = $2
WHERE id = $1
RETURNING *;


  
