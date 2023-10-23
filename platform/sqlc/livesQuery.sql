-- name: CreateLiveStream :one
INSERT INTO lives (userId, title, description, start_time, end_time)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetLiveStreamById :one
SELECT * FROM lives WHERE streamId = $1;

-- name: ListlivesByUserId :many
SELECT * FROM lives WHERE userId = $1 ORDER BY start_time DESC;

-- name: UpdateLiveStreamById :exec
UPDATE lives
SET title = $2, description = $3, start_time = $4, end_time = $5, updated_at = current_timestamp
WHERE streamId = $1;

-- name: DeleteLiveStreamById :exec
DELETE FROM lives WHERE streamId = $1;
