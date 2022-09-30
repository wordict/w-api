build:
	go build -o bin/wordict  ./cmd/wordict

run:
	./bin/wordict

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
	mockgen -destination=internal/mocks/web/mock_service.go -package mocks github.com/eneskzlcn/wordict/internal/web Service
	mockgen -destination=internal/mocks/web/mock_renderer.go -package mocks github.com/eneskzlcn/wordict/internal/web Renderer
	mockgen -destination=internal/mocks/web/mock_logger.go -package mocks github.com/eneskzlcn/wordict/internal/core/logger Logger
	mockgen -destination=internal/mocks/web/mock_session.go -package mocks github.com/eneskzlcn/wordict/internal/core/session Session
	mockgen -destination=internal/core/mocks/mock_zap_logger.go -package mocks github.com/eneskzlcn/wordict/internal/core/logger ZapLogger
