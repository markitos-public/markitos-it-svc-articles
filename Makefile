test-unit:
	go test -count=1 ./internal/domain/...

test-integration:
	go test -count=1 ./internal/infraestructure/...

start-db:
	docker compose -f docker-compose.dev.yaml up -d

stop-db:
	docker compose -f docker-compose.dev.yaml down

start:
	SERVER_ADDRESS=":8081" DATABASE_DSN="host=localhost port=5431 user=postgres password=postgres dbname=faqs_test sslmode=disable" go run cmd/rest/main.go