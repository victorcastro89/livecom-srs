-- Create the "live_stats" table
CREATE TABLE IF NOT EXISTS live_stats (
    live_stat_id SERIAL PRIMARY KEY,
    live_id INT UNIQUE REFERENCES live(live_id) ON DELETE CASCADE, -- Reference to the corresponding live
    status VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp
);