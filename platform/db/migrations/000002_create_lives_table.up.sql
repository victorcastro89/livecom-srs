CREATE EXTENSION IF NOT EXISTS pgcrypto;
-- Livestreams Table
CREATE TABLE IF NOT EXISTS live (
    live_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE, -- A livetream belongs to a User
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    scheduled_start_time TIMESTAMP,
    scheduled_end_time TIMESTAMP,
    live_app_name VARCHAR(255),-- Name of the livetream app
    stream_name VARCHAR(255), -- Name of the livetream
    live_secret_encrypted TEXT,
    live_secret_hash TEXT,
    stream_broadcast_url_encrypted TEXT, -- URL to transmit the livetream
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp
);
-- Index on user_id
CREATE INDEX idx_live_user_id ON live(user_id);

-- Index on start_time and end_time (Composite Index)
CREATE INDEX idx_live_start_end_time ON live(start_time, end_time);

-- Index on title (using a text pattern ops for efficient text searching)
CREATE INDEX idx_live_title ON live(title text_pattern_ops);

CREATE INDEX idx_live_secret_hash ON live(live_secret_hash);

CREATE INDEX idx_live_stream_name ON live(stream_name);
