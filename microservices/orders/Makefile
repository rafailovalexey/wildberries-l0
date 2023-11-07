download:
	go mod download

build:
	go build -o bin/main main.go

generate:
	mockgen -source=internal/api/api.go -destination=internal/api/mocks/mock-api.go
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock-service.go
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock-repository.go
	mockgen -source=internal/converter/converter.go -destination=internal/converter/mocks/mock-converter.go
