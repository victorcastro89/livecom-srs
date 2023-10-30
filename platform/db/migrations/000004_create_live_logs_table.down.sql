-- Drop the index on user_id
DROP INDEX IF EXISTS idx_live_log_client_id;

-- Drop the users table if it exists
DROP TABLE IF EXISTS live_log;
