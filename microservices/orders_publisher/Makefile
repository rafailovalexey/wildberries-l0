download:
	go mod download

build:
	go build -o bin/main main.go

generate:
	mockgen -source=internal/controller/controller.go -destination=internal/controller/mocks/controller_mock.go
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/service_mock.go
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/repository_mock.go

tests:
	go test ./...

.PHONY: download, build, generate, tests
