APP_NAME:=alertflow
VERSION:=$(shell git describe --tags --always --dirty="-dev")
LDFLAGS:=-X main.Version=$(VERSION) -w -s

.PHONY: all

linux:
	GOOS=linux GOARCH=amd64 go build -mod vendor -ldflags="$(LDFLAGS)" -o bin/$(APP_NAME) main.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -mod vendor -ldflags="$(LDFLAGS)" -o bin/$(APP_NAME)-darwin main.go

fmt:
	gofmt -s -w main.go

build: fmt linux darwin

run:
	go run main.go

vendor:
	go mod vendor

all: vendor build
