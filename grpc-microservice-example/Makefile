PROJECT_NAME=grpc-microservice-example
MODULE_NAME=grpc-microservice-example

.DEFAULT_GOAL := server

.PHONY: get
get:
	@go mod tidy
	
.PHONY: server
server:
	@go run ./server/main.go

.PHONY: client
client:
	@go run ./client/main.go	

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html=coverage.out




