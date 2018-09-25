
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:
	go build -o mutant-srv main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t mutant-srv:latest
