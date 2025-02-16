.PHONY: run migrate-up migrate-down

run:
	go run cmd/web/main.go


# Create migration
migration_create:
	@echo "Creating migration..."
	migrate create -ext sql -dir db/migrations/ -seq $(NAME)

# Migrate
migration_up: 
	@echo "Running migration up..."
	migrate -path db/migrations/ -database "postgresql://rizkirmdhn:rizkirmdhn@localhost:5432/segokucing?sslmode=disable" -verbose up

migration_down:
	@echo "Running migration down..."	
	migrate -path db/migrations/ -database "postgresql://rizkirmdhn:rizkirmdhn@localhost:5432/segokucing?sslmode=disable" -verbose down

migration_fix: 
	@echo "Running migration fix..."	
	migrate -path db/migrations/ -database "postgresql://rizkirmdhn:rizkirmdhn@localhost:5432/segokucing?sslmode=disable" force VERSION

# Seed
seed:
	@echo
	@echo "Seeding data..."
	@echo
	go run cmd/seed/main.go
