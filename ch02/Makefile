.DEFAULT_GOAL := build

.PHONY: fmt lint vet build run

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...

build: vet
	go build chapter.go

run: vet
	go run chapter.go
