-- Drop the index on user_id
DROP INDEX idx_live_user_id;

-- Drop the composite index on start_time and end_time
DROP INDEX idx_live_start_end_time;

-- Drop the index on title
DROP INDEX idx_live_title;

-- Drop the users table if it exists
DROP TABLE IF EXISTS live;
