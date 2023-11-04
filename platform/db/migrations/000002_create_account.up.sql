-- account Table
CREATE TABLE IF NOT EXISTS account (
    account_id UUID PRIMARY KEY,
    account_name VARCHAR(255) ,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp
);



