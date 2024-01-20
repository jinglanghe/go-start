PROJECT_NAME=go-start
PKG=github.com/jinglanghe/go-start
GIT_COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS?="-X ${PKG}/cmd/sub_commands.gitCommit=${GIT_COMMIT} -X ${PKG}/cmd/sub_commands.buildDate=${BUILD_DATE}"

get-deps:
	go mod tidy
	go mod download

vet-check-all: get-deps
	go vet ./...

gosec-check-all: get-deps
	gosec ./...

bin: get-deps
	go build -o ${PROJECT_NAME} -ldflags ${LDFLAGS} main.go

gen-swagger:
	swag init -g cmd/${PROJECT_NAME}.go -o api

build: get-deps bin

