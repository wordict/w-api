build:
	go build -o bin/wordict  ./cmd/wordict

start:
	clear && go build -o bin/wordict  ./cmd/wordict && ./bin/wordict

clean:
	rm -rf bin/ && rm -rf internal/mocks/ && rm -rf internal/server/mocks/

test:
	go test ./...

migrate-tables:
	go build -o postgres_migrate  ./seed/cmd  && ./postgres_migrate -type=migrate && rm -rf postgres_migrate && clear

drop-tables:
	go build -o postgres_drop  ./seed/cmd  && ./postgres_drop -type=drop && rm -rf postgres_drop && clear

generate-mocks:
	mockgen -destination=internal/mocks/handler/mock_service.go -package mocks github.com/wordict/w-api/internal/handler Service
	mockgen -destination=internal/mocks/logger/mock_logger.go -package mocks github.com/wordict/w-api/internal/core/logger Logger

