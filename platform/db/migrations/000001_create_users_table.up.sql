CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY,
    firebase_uid  VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    email_verified BOOLEAN,
    first_name VARCHAR(255) ,
    last_name VARCHAR(255) ,
    display_name VARCHAR(255),
    photo_url TEXT,
    phone_number VARCHAR(20) NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
