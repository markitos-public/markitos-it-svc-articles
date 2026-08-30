test-unit:
	go test -count=1 ./internal/domain/...

test-integration:
	go test -count=1 ./internal/infraestructure/...

start-db:
	docker compose -f docker-compose.yaml up -d

stop-db:
	docker compose -f docker-compose.yaml down