migrate:
	migrate create -dir ./internal/storage/database -ext sql db

migrate-up:
	migrate -path ./internal/storage/database -database "postgres://postgres:dilshod@localhost:5432/anyy?sslmode=disable" up

migrate-down:
	migrate -path ./internal/storage/database -database "postgres://postgres:dilshod@localhost:5432/anyy?sslmode=disable" down

migrate-force:
	migrate -path ./internal/storage/database -database "postgres://postgres:dilshod@localhost:5432/anyy?sslmode=disable" force 20240630141811

run-api:
	@go run cmd/main.go

tidy:
	@go mod tidy
run-cmd:
	@go run cmd/main.go