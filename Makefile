build:
	go build -o bin/softdare  ./cmd/softdare

run:
	./bin/softdare

start:
	go build -o bin/softdare  ./cmd/softdare && ./bin/softdare

generate-mocks:
	mockgen -destination=internal/mocks/server/mock_root_handler.go -package mocks github.com/eneskzlcn/softdare/internal/server RootHandler
	mockgen -destination=internal/mocks/server/mock_session_provider.go -package mocks github.com/eneskzlcn/softdare/internal/server Session
	mockgen -destination=internal/mocks/home/mock_home_renderer.go -package mocks github.com/eneskzlcn/softdare/internal/home Renderer
	mockgen -destination=internal/mocks/home/mock_home_service.go -package mocks github.com/eneskzlcn/softdare/internal/home HomeService
	mockgen -destination=internal/mocks/home/mock_session_provider.go -package mocks github.com/eneskzlcn/softdare/internal/home SessionProvider
	mockgen -destination=internal/mocks/login/mock_session_provider.go -package mocks github.com/eneskzlcn/softdare/internal/login SessionProvider
	mockgen -destination=internal/mocks/login/mock_login_renderer.go -package mocks github.com/eneskzlcn/softdare/internal/login Renderer
	mockgen -destination=internal/mocks/login/mock_login_service.go -package mocks github.com/eneskzlcn/softdare/internal/login LoginService
	mockgen -destination=internal/mocks/login/mock_login_repository.go -package mocks github.com/eneskzlcn/softdare/internal/login LoginRepository
	mockgen -destination=internal/mocks/post/mock_post_repository.go -package mocks github.com/eneskzlcn/softdare/internal/post PostRepository
	mockgen -destination=internal/mocks/post/mock_post_service.go -package mocks github.com/eneskzlcn/softdare/internal/post PostService
	mockgen -destination=internal/mocks/post/mock_session_provider.go -package mocks github.com/eneskzlcn/softdare/internal/post SessionProvider

