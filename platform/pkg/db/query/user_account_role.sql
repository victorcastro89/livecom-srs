-- name: CreateUserAccountRoleRelation :one

INSERT INTO user_account_role (user_id, account_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserAccountRoleRelation :exec
DELETE FROM user_account_role
WHERE user_id = $1 AND account_id = $2;


-- name: UpdateUserAccountRoleRelation :one 
UPDATE user_account_role
SET role = $3
WHERE user_id = $1 AND account_id = $2
RETURNING *;


-- name: GetUserAccountAndRoleRelation :many
SELECT user_id, account_id, role
FROM user_account_role
WHERE user_id = $1 ;
