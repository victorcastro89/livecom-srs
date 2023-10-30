-- name: CreateLiveStat
-- :live_id, :viewers_count, :status are placeholders for the respective values
INSERT INTO live_stats (live_id,  status)
VALUES (:live_id,  :status)
RETURNING live_stat_id;

-- name: GetLiveStatByID
SELECT live_stat_id, live_id, status, created_at, updated_at
FROM live_stats
WHERE live_stat_id = :live_stat_id;



-- name: DeleteLiveStat
-- :live_stat_id is a placeholder for the live_stat_id to delete
DELETE FROM live_stats
WHERE live_stat_id = :live_stat_id;
