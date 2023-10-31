-- name: CreateLive :one
INSERT INTO live (user_id, title, description, start_time, end_time, scheduled_start_time, scheduled_end_time, live_app_name, stream_name, live_secret_hash,   live_secret_encrypted, stream_broadcast_url_encrypted)
VALUES (
  @user_id,
  @title,
  @description,
  @start_time,
  @end_time,
  @scheduled_start_time,
  @scheduled_end_time,
  @live_app_name,
  @stream_name,
  @live_secret_hash,
  @live_secret_encrypted,
  @stream_broadcast_url_encrypted
)
RETURNING *;
-- name: GetLiveBySecretHashAppAndStream :one
SELECT live_id, user_id, title, description, start_time, end_time, scheduled_start_time, scheduled_end_time, live_app_name, stream_name,live_secret_hash, live_secret_encrypted, stream_broadcast_url_encrypted, created_at
FROM live
WHERE stream_name = @stream_name AND live_secret_hash =  @live_secret_hash AND live_app_name =  @live_app_name;


-- name: GetLiveByID :one
SELECT live_id, user_id, title, description, start_time, end_time, scheduled_start_time, scheduled_end_time, live_app_name, stream_name,live_secret_hash,live_secret_encrypted, stream_broadcast_url_encrypted,  created_at
FROM live
WHERE live_id = $1;

-- name: GetLiveWithStatusByID :one
SELECT
    l.live_id,
    l.user_id,
    l.title,
    l.description,
    l.start_time,
    l.end_time,
    l.scheduled_start_time,
    l.scheduled_end_time,
    l.live_app_name,
    l.stream_name,
    l.live_secret_hash,
    l.live_secret_encrypted, 
    l.stream_broadcast_url_encrypted, 
    l.created_at,
    l.updated_at,
    ls.status AS live_status
FROM live l
LEFT JOIN live_stats ls ON l.live_id = ls.live_id
WHERE l.live_id = $1;

-- name: GetLivesByUserID :many
SELECT * FROM live WHERE user_id = $1 ORDER BY start_time DESC;

-- name: GetOngoingLives :many
SELECT * FROM live WHERE start_time <= CURRENT_TIMESTAMP AND (end_time IS NULL OR end_time > CURRENT_TIMESTAMP) ORDER BY start_time DESC;

-- name: GetLiveWithUserDetails :one
SELECT     l.live_id,
    l.user_id,
    l.title,
    l.description,
    l.start_time,
    l.end_time,
    l.scheduled_start_time,
    l.scheduled_end_time,
    l.live_app_name,
    l.stream_name,
    l.live_secret_hash,
    l.live_secret_encrypted, 
    l.stream_broadcast_url_encrypted, 
    l.created_at,
    l.updated_at,
     u.first_name, 
     u.last_name, 
     u.email 
FROM live l 
JOIN users u ON l.user_id = u.user_id 
WHERE l.live_id = $1;


-- name: UpdateLive :one
UPDATE live
SET
  user_id = $1,
  title = $2,
  description = $3,
  start_time = $4,
  end_time = $5,
  scheduled_start_time = $6,
  scheduled_end_time = $7,
  live_app_name = $8,
  stream_name = $9,
  live_secret_hash = $10,
  live_secret_encrypted = $11,
  stream_broadcast_url_encrypted = $12,
  updated_at = current_timestamp
WHERE live_id = $13
RETURNING *;

-- name: DeleteLive :exec
DELETE FROM live
WHERE live_id = $1;

