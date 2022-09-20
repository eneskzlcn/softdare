build:
	go build -o bin/softdare  ./cmd/softdare

run:
	./bin/softdare

start:
	clear && go build -o bin/softdare  ./cmd/softdare && ./bin/softdare

clean:
	rm -rf bin/ && rm -rf internal/mocks/ && rm -rf internal/server/mocks/

test:
	go test ./...

migrate-tables:
	go build ./postgres/migration && ./schema -type=migrate && rm -rf migration && clear

drop-tables:
	go build ./postgres/migration && ./schema -type=drop && rm -rf migration && clear

generate-mocks:
	mockgen -destination=internal/mocks/web/mock_service.go -package mocks github.com/eneskzlcn/softdare/internal/web Service
	mockgen -destination=internal/mocks/web/mock_renderer.go -package mocks github.com/eneskzlcn/softdare/internal/web Renderer
	mockgen -destination=internal/mocks/web/mock_logger.go -package mocks github.com/eneskzlcn/softdare/internal/core/logger Logger
	mockgen -destination=internal/mocks/web/mock_session.go -package mocks github.com/eneskzlcn/softdare/internal/core/session Session
	mockgen -destination=internal/core/mocks/mock_zap_logger.go -package mocks github.com/eneskzlcn/softdare/internal/core/logger ZapLogger
