GOPATH:=$(shell go env GOPATH)

.PHONY: run build test docker-build docker-push

run:
	go run main.go

build:
	go build -o gateway-srv main.go

test:
	go test -v ./... -cover

docker-build:
	docker build -t rodrigodmd/ml-mutant-srv:gateway .

docker-push: docker-build
	docker push rodrigodmd/ml-mutant-srv:gateway
