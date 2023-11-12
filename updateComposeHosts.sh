#!/bin/bash

# Function to update /etc/hosts if an entry does not exist
add_host_entry() {
    local ip=$1
    local hostnames=$2
    if ! grep -qE "^$ip\s+.*$hostnames" /etc/hosts; then
        echo "$ip $hostnames" >> /etc/hosts
    fi
}

# Add your custom host entries
add_host_entry "127.0.0.2" "stream-high.local stream-high"
add_host_entry "127.0.0.3" "stream-low.local stream-low"

# Execute the main command
exec ./objs/srs -c conf/docker.conf
