-- Create the "live_stats" table
CREATE TABLE IF NOT EXISTS live_log (
    live_log_id SERIAL PRIMARY KEY,
    live_id INT REFERENCES live(live_id) ON DELETE CASCADE, -- Reference to the corresponding live
    action VARCHAR(255),
    client_id VARCHAR(255), --Client ID, each connection is a new client ID
    stream_id VARCHAR(255), --Stream ID, each StreamURL is a new client ID
    server_id VARCHAR(255), --Server ID, each SRS Server is a new ID, It changes after reboot
    service_id VARCHAR(255),
    ip VARCHAR(255),
    vhost VARCHAR(255),
    app VARCHAR(255),
    tcUrl VARCHAR(255),
    stream_url_param VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp
 );

CREATE INDEX idx_live_log_client_id ON live_log(client_id);


