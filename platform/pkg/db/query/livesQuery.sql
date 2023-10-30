-- name: CreateLive :one
INSERT INTO live (user_id, title, description, start_time, end_time, scheduled_start_time, scheduled_end_time, live_app_name, stream_name, live_secret, stream_broadcast_url)
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
  pgp_sym_encrypt(@live_secret, @encryptionKey::text),
  pgp_sym_encrypt(@stream_broadcast_url, @encryptionKey::text)
)
RETURNING live_id, user_id, title, description, start_time, end_time, scheduled_start_time, scheduled_end_time, live_app_name, stream_name, pgp_sym_decrypt(live_secret,@encryptionKey::text) as live_secret, pgp_sym_decrypt(stream_broadcast_url,@encryptionKey::text) as stream_broadcast_url, created_at;


-- name: GetLiveByID :one
SELECT live_id, user_id, title, description, start_time, end_time, scheduled_start_time, scheduled_end_time, live_app_name, stream_name, pgp_sym_decrypt(live_secret,@encryptionKey::text) as live_secret, pgp_sym_decrypt(stream_broadcast_url,@encryptionKey::text) as stream_broadcast_url, created_at
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
pgp_sym_decrypt(live_secret,@encryptionKey::text) as live_secret, 
pgp_sym_decrypt(stream_broadcast_url,@encryptionKey::text) as stream_broadcast_url,
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
pgp_sym_decrypt(live_secret,@encryptionKey::text) as live_secret, 
pgp_sym_decrypt(stream_broadcast_url,@encryptionKey::text) as stream_broadcast_url,
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
  live_secret = pgp_sym_encrypt(@live_secret, @encryptionKey::text),
  stream_broadcast_url= pgp_sym_encrypt(@stream_broadcast_url, @encryptionKey::text),
  updated_at = current_timestamp
WHERE live_id = $10
RETURNING *;

-- name: DeleteLive :exec
DELETE FROM live
WHERE live_id = $1;

