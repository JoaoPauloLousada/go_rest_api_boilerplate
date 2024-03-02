DC_LOCAL_FILE = ./build/local/docker-compose.yml

local-up:
	docker compose -f $(DC_LOCAL_FILE) up -d

local-down:
	docker compose -f $(DC_LOCAL_FILE) down

local-logs:
	docker compose -f $(DC_LOCAL_FILE) logs -f

local-run:
	go run ./cmd/api

local-inspect-shema:
	atlas schema inspect -u "postgres://exampleuser:exampleuser@0.0.0.0:5432/exampledb?search_path=public&sslmode=disable" --format "{{ sql . }}"

migration:
	atlas migrate diff $(name) \
  --dir "file://internal/infrastructure/db/migrations" \
  --to "file://internal/infrastructure/db/schema.sql" \
  --dev-url "postgres://exampleuser:exampleuser@0.0.0.0:5432/exampledb?search_path=public&sslmode=disable"

local-migrate:
	atlas migrate apply \
	--dir "file://internal/infrastructure/db/migrations" \
	--url "postgres://exampleuser:exampleuser@0.0.0.0:5432/exampledb?search_path=public&sslmode=disable"