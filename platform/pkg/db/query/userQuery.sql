-- name: CreateUser :one
INSERT INTO users (user_id, email, first_name, last_name, phone_number) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;



-- name: UpdateUser :one
UPDATE users 
SET email = $2, first_name = $3, last_name = $4, phone_number = $5, updated_at = current_timestamp 
WHERE user_id = $1 
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;
