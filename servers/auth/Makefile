.PHONY: migrate_up migrate_down sqlc_generate

# ==============================================================================
# Go migrate postgresql

migrate_force:
	migrate -path pkg/database/migration -database "postgresql://admin:admin@localhost:5432/beta-db?sslmode=disable" -verbose force 1

migrate_up:
	migrate -path pkg/database/migration -database "postgresql://admin:admin@localhost:5432/beta-db?sslmode=disable" -verbose up

migrate_down:
	migrate -path pkg/database/migration -database "postgresql://admin:admin@localhost:5432/beta-db?sslmode=disable" -verbose down

# ==============================================================================
# Go sqlc postgresql

generate:
	sqlc generate
