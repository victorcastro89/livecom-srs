-- name: CreateUser :one
INSERT INTO users (username, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;


-- name: UpdateUserByID :exec
UPDATE users
SET username = $2, email = $3, first_name = $4, last_name = $5, updated_at = current_timestamp
WHERE id = $1
RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;
