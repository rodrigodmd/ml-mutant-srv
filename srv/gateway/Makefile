GOPATH:=$(shell go env GOPATH)

.PHONY: build test docker-build docker-push

build:
	go build -o gateway-srv main.go plugin.go

test:
	go test -v ./... -cover

docker-build:
	docker build -t rodrigodmd/ml-mutant-srv:gateway .

docker-push: docker-build
	docker push rodrigodmd/ml-mutant-srv:gateway
