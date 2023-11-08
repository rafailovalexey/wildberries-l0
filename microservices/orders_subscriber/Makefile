download:
	go mod download

build:
	go build -o bin/main main.go

generate:
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/repository_mock.go
	mockgen -source=internal/converter/converter.go -destination=internal/converter/mocks/converter_mock.go

tests:
	go test -v ./...

.PHONY: download, build, generate, tests
