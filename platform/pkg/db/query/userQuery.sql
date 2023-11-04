-- name: CreateUser :one
INSERT INTO users (user_id, email, email_verified ,firebase_uid, display_name, photo_url, first_name, last_name, phone_number) 
VALUES ($1, $2, $3, $4, $5 ,$6,$7,$8,$9) 
RETURNING *;

-- name: GetUserWithRoleAndAccountByID :one
SELECT
    u.user_id,
    u.firebase_uid,
    u.email,
    u.email_verified,
    u.first_name,
    u.last_name,
    u.display_name,
    u.photo_url,
    u.phone_number,
    ua.role,
    a.account_id,
    a.account_name
FROM
    users u
JOIN
    user_account_role ua ON u.user_id = ua.user_id
JOIN
    account a ON ua.account_id = a.account_id
WHERE
    u.user_id = $1;


-- name: GetUserWithRoleAndAccountByFirebaseUID :one
SELECT
    u.user_id,
    u.firebase_uid,
    u.email,
    u.email_verified,
    u.first_name,
    u.last_name,
    u.display_name,
    u.photo_url,
    u.phone_number,
    u.created_at,
    u.updated_at,
    ua.role,
    a.account_id,
    a.account_name
FROM
    users u
JOIN
    user_account_role ua ON u.user_id = ua.user_id
JOIN
    account a ON ua.account_id = a.account_id
WHERE
    u.firebase_uid = $1;


-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;


-- name: GetUserByFirebaseUID :one
SELECT * FROM users WHERE firebase_uid = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;



-- name: UpdateUser :one
UPDATE users 
SET email = $2, first_name = $3, last_name = $4, display_name =$5,photo_url=$6, phone_number = $5, updated_at = current_timestamp 
WHERE user_id = $1 
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;
