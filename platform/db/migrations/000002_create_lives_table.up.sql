-- Livestreams Table
CREATE TABLE IF NOT EXISTS lives (
    live_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE, -- A livestream belongs to a User
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    stream_url VARCHAR(1000), -- URL to access the livestream
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Index on user_id
CREATE INDEX idx_lives_user_id ON lives(user_id);

-- Index on start_time and end_time (Composite Index)
CREATE INDEX idx_lives_start_end_time ON lives(start_time, end_time);

-- Index on title (using a text pattern ops for efficient text searching)
CREATE INDEX idx_lives_title ON lives(title text_pattern_ops);