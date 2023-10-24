-- name: CreateLive :one
INSERT INTO lives (live_id, user_id, title, description, start_time, end_time, stream_url) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetLiveByID :one
SELECT * FROM lives WHERE live_id = $1;

-- name: GetLivesByUserID :many
SELECT * FROM lives WHERE user_id = $1 ORDER BY start_time DESC;

-- name: GetOngoingLives :many
SELECT * FROM lives WHERE start_time <= CURRENT_TIMESTAMP AND (end_time IS NULL OR end_time > CURRENT_TIMESTAMP) ORDER BY start_time DESC;

-- name: GetLiveWithUserDetails :one
SELECT l.*, u.first_name, u.last_name, u.email 
FROM lives l 
JOIN users u ON l.user_id = u.user_id 
WHERE l.live_id = $1;


-- name: UpdateLive :one
UPDATE lives 
SET user_id = $2, title = $3, description = $4, start_time = $5, end_time = $6, stream_url = $7 
WHERE live_id = $1 
RETURNING *;

-- name: DeleteLive :exec
DELETE FROM lives WHERE live_id = $1;

