#!/bin/bash

# Start docker containers
docker-compose up -d

# Wait for PostgreSQL to start
echo "Waiting for PostgreSQL to start..."
sleep 10

# Run database migrations
sh ./migratedb.sh

