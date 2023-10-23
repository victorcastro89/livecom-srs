migrate -database "postgres://livecom:livecom@localhost:5432/livecom?sslmode=disable" -path db/migrations up
# Print success message
echo "Migration completed successfully!"