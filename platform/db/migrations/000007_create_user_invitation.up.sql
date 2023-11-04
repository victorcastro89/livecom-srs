-- Define ENUM type for roles
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'invite_status_type') THEN
        CREATE TYPE invite_status_type AS ENUM('pending', 'accepted');
    END IF;
END$$;

-- account Table
CREATE TABLE IF NOT EXISTS user_invite (
    invite_id UUID PRIMARY KEY ,
    invited_by UUID NOT NULL,
    account_id UUID NOT NULL,
    email_invited VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL
);



