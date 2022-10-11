init:
	go mod vendor

run:
	go run cmd/cms/main.go

test:
	go test ./internal/control/...

build:
	docker build -t upday-task-fe:latest .
