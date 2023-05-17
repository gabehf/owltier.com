.PHONY: client

client:
	@yarn --cwd ./client vite

api:
	@echo 'Starting API Server...'
	@go run ./main.go

test: test.api

test.api:
	go test -v ./...

dynamo.start:
	@cd ./dev/db && docker compose up -d

dynamo.stop:
	@cd ./dev/db && docker compose down

dynamo.list:
	@aws dynamodb list-tables --endpoint-url http://localhost:8000

dynamo.scan:
	@aws dynamodb scan --table-name owltier-local --endpoint-url http://localhost:8000

dynamo.reset:
	@cd ./dev/scripts && ./delete-table.sh
	@sleep 1
	@cd ./dev/scripts && ./create-table.sh