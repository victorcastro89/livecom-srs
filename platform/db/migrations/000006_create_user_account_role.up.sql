-- Define ENUM type for roles
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_type') THEN
        CREATE TYPE role_type AS ENUM('owner', 'admin', 'cohost');
    END IF;
END$$;


-- User account Table with an ENUM type for roles
CREATE TABLE IF NOT EXISTS user_account_role (
    user_id UUID NOT NULL,
    account_id UUID NOT NULL,
    role role_type NOT NULL,
    PRIMARY KEY (user_id, account_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES account (account_id) ON DELETE CASCADE
);

-- Indexes for the User account Table to improve query performance
CREATE INDEX IF NOT EXISTS idx_user_account_role_user ON user_account_role(user_id);
CREATE INDEX IF NOT EXISTS idx_user_account_role_account ON user_account_role(account_id);
