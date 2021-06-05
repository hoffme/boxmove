SHELL:=/bin/bash -O extglob
BINARY=builds/boxmove
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

#go tool commands
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/boxmove/main.go

run:
	@go run cmd/boxmove/main.go

## protobuf
protobuf:
	protoc -I=./proto --go_out=. --go-grpc_out=. \
		boxes.proto \
		clients.proto \
		moves.proto

## tests
test:
	@go test ./...

## docker
buildImage:
	docker build -t boxmove . && docker image prune --filter label=stage=builder

## docker compose
up:
	docker-compose up --build

down:
	docker-compose down --remove-orphans