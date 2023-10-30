-- name: CreateLiveLog
-- :live_id, :action, :client_id, :stream_id, :server_id, :service_id, :ip, :vhost, :app, :tcUrl, :stream_url_param are placeholders for the respective values
INSERT INTO live_log (live_id, action, client_id, stream_id, server_id, service_id, ip, vhost, app, tcUrl, stream_url_param)
VALUES (:live_id, :action, :client_id, :stream_id, :server_id, :service_id, :ip, :vhost, :app, :tcUrl, :stream_url_param)
RETURNING live_log_id;

-- name: GetLiveLogByID
SELECT live_log_id, live_id, action, client_id, stream_id, server_id, service_id, ip, vhost, app, tcUrl, stream_url_param, created_at, updated_at
FROM live_log
WHERE live_log_id = :live_log_id;

-- -- name: UpdateLiveLog
-- -- :live_log_id is a placeholder for the live_log_id to update
-- UPDATE live_log
-- SET
--  action = :action,
--  client_id = :client_id,
--  stream_id = :stream_id,
--  server_id = :server_id,
--  service_id = :service_id,
--  ip = :ip,
--  vhost = :vhost,
--  app = :app,
--  tcUrl = :tcUrl,
--  stream_url_param = :stream_url_param,
--  updated_at = CURRENT_TIMESTAMP
-- WHERE live_log_id = :live_log_id;

-- name: DeleteLiveLog
-- :live_log_id is a placeholder for the live_log_id to delete
DELETE FROM live_log
WHERE live_log_id = :live_log_id;
