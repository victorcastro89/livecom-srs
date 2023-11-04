-- name: CreateAccount :one
INSERT INTO account (account_id, account_name)
VALUES ($1, $2)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM account WHERE account_id= $1;


-- name: GetAccountsAndRolesByUserID :many
SELECT  uar.account_id, ac.account_name,uar.user_id, uar.role
FROM user_account_role uar
JOIN account ac ON uar.account_id = ac.account_id
WHERE uar.user_id = $1 
ORDER BY ac.created_at DESC;
