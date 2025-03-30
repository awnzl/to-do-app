DB_URL="postgres://test:test@db:5432/todoapp-db?sslmode=disable"

.PHONY: all up down logs migrate-up migrate-down clean

# Main commands
init: up migrate-up

# Docker commands
up:
	docker compose up -d --build

down:
	docker compose down

rmdb:
	docker volume rm todoapp_postgres_data

logs:
	docker compose logs -f

# Database commands (executed in app container)
migrate-up:
	docker compose exec todoapp migrate -path migrations -database ${DB_URL} up

migrate-down:
	docker compose exec todoapp migrate -path migrations -database ${DB_URL} down

# Clean temporary files
clean:
	rm -rf backend/bin frontend/dist frontend/node_modules
