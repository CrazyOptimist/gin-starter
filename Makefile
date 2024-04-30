test:
	go test -v $$(go list ./... | grep -v /docs)
vet:
	go vet -v $$(go list ./... | grep -v /docs)
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/server
migrate:
	go run ./cmd/cli/main.go --migrate
docs:
	rm -rf docs/* && swag init -d ./cmd/server,./
docker:
	docker compose -f ./deployments/compose.yml build
up:
	docker compose -f ./deployments/compose.yml up -d
down:
	docker compose -f ./deployments/compose.yml down
log:
	docker compose -f ./deployments/compose.yml logs -f

.PHONY: docs test build docker log
