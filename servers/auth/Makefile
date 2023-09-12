.PHONY: migrate_up migrate_down sqlc

# ==============================================================================
# Go migrate postgresql

migrate_force:
	migrate -path pkg/database/migration -database "postgresql://admin:admin@192.168.1.100:5432/beta-db?sslmode=disable" -verbose force 1

migrate_up:
	migrate -path pkg/database/migration -database "postgresql://admin:admin@192.168.1.100:5432/beta-db?sslmode=disable" -verbose up

migrate_down:
	migrate -path pkg/database/migration -database "postgresql://admin:admin@192.168.1.100:5432/beta-db?sslmode=disable" -verbose down

# ==============================================================================
# Go sqlc postgresql

sqlc:
	sqlc generate
