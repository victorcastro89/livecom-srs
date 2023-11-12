 migrate create -ext sql -dir db/migrations -seq create_live_logs_table    

 migrate -database "postgres://livecom:livecom@localhost:5432/livecom?sslmode=disable" -path db/migrations up

 sqlc generate