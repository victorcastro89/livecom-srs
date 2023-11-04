-- Start transaction
BEGIN;

-- Drop the user_account_roles table
DROP TABLE IF EXISTS user_account_role;


-- Drop the role_type ENUM (caution: ensure that no other table is using this type before dropping it)
DROP TYPE IF EXISTS role_type;

-- Drop the indexes (they should be dropped automatically with the tables, but we include them for completeness)
DROP INDEX IF EXISTS idx_user_account_roles_user;
DROP INDEX IF EXISTS idx_user_account_roles_account;


-- End transaction
COMMIT;
