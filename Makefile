.PHONY: build run test clean swag

build:
	go build -o bin/simpleorder ./cmd/app

run:
	go run cmd/app/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

swag:
	swag init -g cmd/app/main.go
